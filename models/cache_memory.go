package models

import (
	"container/list"
	"encoding/json"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
)

// 缓存项
type CacheItem struct {
	Key        string      `json:"key"`
	Value      interface{} `json:"value"`
	CreatedAt  time.Time   `json:"created_at"`
	AccessedAt time.Time   `json:"accessed_at"`
	AccessCount int        `json:"access_count"`
	TTL        time.Duration `json:"ttl"`
	element    *list.Element // LRU链表元素
}

// 内存缓存管理器
type MemoryCache struct {
	cache    map[string]*CacheItem
	lruList  *list.List
	maxSize  int
	defaultTTL time.Duration
	mutex    sync.RWMutex
	stats    *MemoryCacheStats
	cleaner  *time.Ticker
}

// 内存缓存统计信息
type MemoryCacheStats struct {
	Size         int           `json:"size"`
	MaxSize      int           `json:"max_size"`
	Hits         int64         `json:"hits"`
	Misses       int64         `json:"misses"`
	Evictions    int64         `json:"evictions"`
	Expirations  int64         `json:"expirations"`
	HitRate      float64       `json:"hit_rate"`
	TTL          time.Duration `json:"ttl"`
	LastCleanup  time.Time     `json:"last_cleanup"`
}

// 创建内存缓存
func NewMemoryCache(maxSize int, defaultTTL time.Duration) *MemoryCache {
	cache := &MemoryCache{
		cache:      make(map[string]*CacheItem),
		lruList:    list.New(),
		maxSize:    maxSize,
		defaultTTL: defaultTTL,
		stats: &MemoryCacheStats{
			MaxSize: maxSize,
			TTL:     defaultTTL,
		},
	}

	// 启动清理定时器
	cache.startCleaner()

	return cache
}

// 设置缓存项
func (mc *MemoryCache) Set(key string, value interface{}) {
	mc.SetWithTTL(key, value, mc.defaultTTL)
}

// 设置缓存项（带TTL）
func (mc *MemoryCache) SetWithTTL(key string, value interface{}, ttl time.Duration) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	now := time.Now()

	// 检查是否已存在
	if existingItem, exists := mc.cache[key]; exists {
		// 更新现有项
		existingItem.Value = value
		existingItem.AccessedAt = now
		existingItem.TTL = ttl
		existingItem.AccessCount++
		
		// 移动到LRU链表头部
		mc.lruList.MoveToFront(existingItem.element)
		return
	}

	// 检查缓存大小限制
	if len(mc.cache) >= mc.maxSize {
		mc.evictLRU()
	}

	// 创建新缓存项
	item := &CacheItem{
		Key:         key,
		Value:       value,
		CreatedAt:   now,
		AccessedAt:  now,
		AccessCount: 1,
		TTL:         ttl,
	}

	// 添加到LRU链表头部
	item.element = mc.lruList.PushFront(item)
	mc.cache[key] = item

	mc.stats.Size = len(mc.cache)
}

// 获取缓存项
func (mc *MemoryCache) Get(key string) (interface{}, bool) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	item, exists := mc.cache[key]
	if !exists {
		mc.stats.Misses++
		mc.updateHitRate()
		return nil, false
	}

	// 检查是否过期
	if mc.isExpired(item) {
		mc.deleteItem(key, item)
		mc.stats.Misses++
		mc.stats.Expirations++
		mc.updateHitRate()
		return nil, false
	}

	// 更新访问信息
	item.AccessedAt = time.Now()
	item.AccessCount++

	// 移动到LRU链表头部
	mc.lruList.MoveToFront(item.element)

	mc.stats.Hits++
	mc.updateHitRate()

	return item.Value, true
}

// 删除缓存项
func (mc *MemoryCache) Delete(key string) bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	item, exists := mc.cache[key]
	if !exists {
		return false
	}

	mc.deleteItem(key, item)
	return true
}

// 检查缓存项是否存在
func (mc *MemoryCache) Exists(key string) bool {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	item, exists := mc.cache[key]
	if !exists {
		return false
	}

	return !mc.isExpired(item)
}

// 获取缓存项（不更新访问时间）
func (mc *MemoryCache) Peek(key string) (interface{}, bool) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	item, exists := mc.cache[key]
	if !exists {
		return nil, false
	}

	if mc.isExpired(item) {
		return nil, false
	}

	return item.Value, true
}

// 清空缓存
func (mc *MemoryCache) Clear() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mc.cache = make(map[string]*CacheItem)
	mc.lruList = list.New()
	mc.stats.Size = 0
	mc.stats.Evictions = 0
	mc.stats.Expirations = 0

	logs.Info("[MemoryCache] 缓存已清空")
}

// 获取所有键
func (mc *MemoryCache) Keys() []string {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	keys := make([]string, 0, len(mc.cache))
	for key, item := range mc.cache {
		if !mc.isExpired(item) {
			keys = append(keys, key)
		}
	}

	return keys
}

// 获取缓存大小
func (mc *MemoryCache) Size() int {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()
	return len(mc.cache)
}

// 获取统计信息
func (mc *MemoryCache) GetStats() *MemoryCacheStats {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	// 复制统计信息
	stats := *mc.stats
	stats.Size = len(mc.cache)
	return &stats
}

// 重置统计信息
func (mc *MemoryCache) ResetStats() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	mc.stats.Hits = 0
	mc.stats.Misses = 0
	mc.stats.Evictions = 0
	mc.stats.Expirations = 0
	mc.stats.HitRate = 0
}

// LRU淘汰
func (mc *MemoryCache) evictLRU() {
	if mc.lruList.Len() == 0 {
		return
	}

	// 获取最后一个元素（最少使用）
	element := mc.lruList.Back()
	if element == nil {
		return
	}

	item := element.Value.(*CacheItem)
	mc.deleteItem(item.Key, item)
	mc.stats.Evictions++

	logs.Debug("[MemoryCache] LRU淘汰缓存项: %s", item.Key)
}

// 删除缓存项（内部方法）
func (mc *MemoryCache) deleteItem(key string, item *CacheItem) {
	delete(mc.cache, key)
	if item.element != nil {
		mc.lruList.Remove(item.element)
	}
	mc.stats.Size = len(mc.cache)
}

// 检查是否过期
func (mc *MemoryCache) isExpired(item *CacheItem) bool {
	if item.TTL <= 0 {
		return false // 永不过期
	}
	return time.Since(item.CreatedAt) > item.TTL
}

// 更新命中率
func (mc *MemoryCache) updateHitRate() {
	total := mc.stats.Hits + mc.stats.Misses
	if total > 0 {
		mc.stats.HitRate = float64(mc.stats.Hits) / float64(total) * 100
	}
}

// 启动清理定时器
func (mc *MemoryCache) startCleaner() {
	cleanupInterval := 5 * time.Minute
	if GlobalConfigManager != nil && GlobalConfigManager.GetCacheConfig() != nil {
		cleanupInterval = GlobalConfigManager.GetCacheConfig().CleanupInterval
	}

	mc.cleaner = time.NewTicker(cleanupInterval)
	go func() {
		for range mc.cleaner.C {
			mc.cleanup()
		}
	}()
}

// 清理过期项
func (mc *MemoryCache) cleanup() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	expiredKeys := make([]string, 0)
	
	for key, item := range mc.cache {
		if mc.isExpired(item) {
			expiredKeys = append(expiredKeys, key)
		}
	}

	// 删除过期项
	for _, key := range expiredKeys {
		if item, exists := mc.cache[key]; exists {
			mc.deleteItem(key, item)
			mc.stats.Expirations++
		}
	}

	mc.stats.LastCleanup = time.Now()

	if len(expiredKeys) > 0 {
		logs.Debug("[MemoryCache] 清理过期缓存项 %d 个", len(expiredKeys))
	}
}

// 停止缓存
func (mc *MemoryCache) Stop() {
	if mc.cleaner != nil {
		mc.cleaner.Stop()
	}
	logs.Info("[MemoryCache] 内存缓存已停止")
}

// 获取缓存项详情
func (mc *MemoryCache) GetItemDetails(key string) (*CacheItem, bool) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	item, exists := mc.cache[key]
	if !exists {
		return nil, false
	}

	if mc.isExpired(item) {
		return nil, false
	}

	// 复制缓存项（不包含element）
	itemCopy := &CacheItem{
		Key:         item.Key,
		Value:       item.Value,
		CreatedAt:   item.CreatedAt,
		AccessedAt:  item.AccessedAt,
		AccessCount: item.AccessCount,
		TTL:         item.TTL,
	}

	return itemCopy, true
}

// 批量设置
func (mc *MemoryCache) SetBatch(items map[string]interface{}) {
	for key, value := range items {
		mc.Set(key, value)
	}
}

// 批量获取
func (mc *MemoryCache) GetBatch(keys []string) map[string]interface{} {
	result := make(map[string]interface{})
	for _, key := range keys {
		if value, exists := mc.Get(key); exists {
			result[key] = value
		}
	}
	return result
}

// 批量删除
func (mc *MemoryCache) DeleteBatch(keys []string) int {
	deleted := 0
	for _, key := range keys {
		if mc.Delete(key) {
			deleted++
		}
	}
	return deleted
}

// 设置TTL
func (mc *MemoryCache) SetTTL(key string, ttl time.Duration) bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	item, exists := mc.cache[key]
	if !exists {
		return false
	}

	item.TTL = ttl
	return true
}

// 获取TTL
func (mc *MemoryCache) GetTTL(key string) (time.Duration, bool) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	item, exists := mc.cache[key]
	if !exists {
		return 0, false
	}

	if mc.isExpired(item) {
		return 0, false
	}

	remaining := item.TTL - time.Since(item.CreatedAt)
	if remaining < 0 {
		remaining = 0
	}

	return remaining, true
}

// 导出缓存状态
func (mc *MemoryCache) ExportState() (string, error) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	state := make(map[string]interface{})
	state["stats"] = mc.stats
	state["size"] = len(mc.cache)
	state["max_size"] = mc.maxSize
	state["default_ttl"] = mc.defaultTTL

	// 导出所有有效的缓存项（不包含值，只包含元数据）
	items := make(map[string]interface{})
	for key, item := range mc.cache {
		if !mc.isExpired(item) {
			items[key] = map[string]interface{}{
				"created_at":   item.CreatedAt,
				"accessed_at":  item.AccessedAt,
				"access_count": item.AccessCount,
				"ttl":          item.TTL,
			}
		}
	}
	state["items"] = items

	data, err := json.Marshal(state)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// 获取热门缓存项
func (mc *MemoryCache) GetTopItems(limit int) []*CacheItem {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	items := make([]*CacheItem, 0, len(mc.cache))
	for _, item := range mc.cache {
		if !mc.isExpired(item) {
			itemCopy := &CacheItem{
				Key:         item.Key,
				CreatedAt:   item.CreatedAt,
				AccessedAt:  item.AccessedAt,
				AccessCount: item.AccessCount,
				TTL:         item.TTL,
			}
			items = append(items, itemCopy)
		}
	}

	// 按访问次数排序
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			if items[i].AccessCount < items[j].AccessCount {
				items[i], items[j] = items[j], items[i]
			}
		}
	}

	if limit > 0 && limit < len(items) {
		items = items[:limit]
	}

	return items
}