package models

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
)

// 缓存的告警信息
type CachedAlert struct {
	Fingerprint   *AlertFingerprint `json:"fingerprint"`
	FirstSeen     time.Time         `json:"first_seen"`
	LastSeen      time.Time         `json:"last_seen"`
	Count         int               `json:"count"`
	Status        string            `json:"status"`
	LastAlert     *StandardAlert    `json:"last_alert"`
	SuppressUntil time.Time         `json:"suppress_until"`
}

// 去重结果
type DeduplicationResult struct {
	ShouldSend bool         `json:"should_send"`
	Action     string       `json:"action"` // new, duplicate, suppressed, aggregated
	Count      int          `json:"count"`
	Cached     *CachedAlert `json:"cached"`
	Reason     string       `json:"reason"`
}

// 告警去重管理器
type AlertDeduplicator struct {
	config       *DeduplicationConfig
	fingerprinter *AlertFingerprinter
	cache        map[string]*CachedAlert // 指纹 -> 缓存告警
	mutex        sync.RWMutex
	cleaner      *time.Ticker // 清理定时器
	stats        *DeduplicationStats
	statsMutex   sync.RWMutex
}

// 创建告警去重管理器
func NewAlertDeduplicator(config *DeduplicationConfig, fingerprintConfig *FingerprintConfig) *AlertDeduplicator {
	if config == nil {
		config = &DeduplicationConfig{
			Enabled:         true,
			TimeWindow:      5 * time.Minute,
			MaxCount:        5,
			SuppressResolved: true,
			GroupByLabels:   []string{"alertname", "instance", "severity"},
			Policy:          "strict",
		}
	}

	deduplicator := &AlertDeduplicator{
		config:        config,
		fingerprinter: NewAlertFingerprinter(fingerprintConfig),
		cache:         make(map[string]*CachedAlert),
		stats: &DeduplicationStats{
			TotalRecords:    0,
			ActiveRecords:   0,
			TodayRecords:    0,
			TotalDuplicates: 0,
		},
	}

	// 启动清理定时器
	deduplicator.startCleaner()

	return deduplicator
}

// 检查是否应该发送告警
func (ad *AlertDeduplicator) ShouldSend(alert *StandardAlert) (*DeduplicationResult, error) {
	if !ad.config.Enabled {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "disabled",
			Count:      1,
			Reason:     "去重功能已禁用",
		}, nil
	}

	// 生成指纹
	fingerprint := ad.fingerprinter.GenerateFingerprint(alert)
	
	ad.mutex.Lock()
	defer ad.mutex.Unlock()

	cached, exists := ad.cache[fingerprint.Hash]

	if !exists {
		// 首次出现的告警
		cached = &CachedAlert{
			Fingerprint: fingerprint,
			FirstSeen:   time.Now(),
			LastSeen:    time.Now(),
			Count:       1,
			Status:      alert.Status,
			LastAlert:   alert,
		}
		ad.cache[fingerprint.Hash] = cached

		// 更新统计
		ad.updateStats("new", 1)

		// 持久化到数据库
		go ad.persistToDatabase(cached)

		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "new",
			Count:      1,
			Cached:     cached,
			Reason:     "首次出现的告警",
		}, nil
	}

	// 检查时间窗口
	if time.Since(cached.LastSeen) > ad.config.TimeWindow {
		// 超出时间窗口，重置计数
		logs.Info("[Deduplicator] 告警超出时间窗口，重置计数: %s", fingerprint.Hash)
		cached.Count = 1
		cached.FirstSeen = time.Now()
		cached.SuppressUntil = time.Time{}
	} else {
		cached.Count++
	}

	cached.LastSeen = time.Now()
	cached.LastAlert = alert
	cached.Status = alert.Status

	// 更新统计
	ad.updateStats("duplicate", 1)

	// 判断是否应该发送
	result := ad.shouldSendBasedOnPolicy(cached, alert)
	
	// 异步更新数据库
	go ad.updateDatabase(cached)

	return result, nil
}

// 基于策略判断是否应该发送
func (ad *AlertDeduplicator) shouldSendBasedOnPolicy(cached *CachedAlert, alert *StandardAlert) *DeduplicationResult {
	// 检查是否被抑制
	if !cached.SuppressUntil.IsZero() && time.Now().Before(cached.SuppressUntil) {
		return &DeduplicationResult{
			ShouldSend: false,
			Action:     "suppressed",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     fmt.Sprintf("告警被抑制到 %s", cached.SuppressUntil.Format("2006-01-02 15:04:05")),
		}
	}

	// 检查恢复告警抑制
	if ad.config.SuppressResolved && alert.IsResolved() && cached.LastAlert != nil && cached.LastAlert.IsFiring() {
		logs.Info("[Deduplicator] 抑制恢复告警: %s", cached.Fingerprint.Hash)
		return &DeduplicationResult{
			ShouldSend: false,
			Action:     "resolved_suppressed",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     "恢复告警被抑制",
		}
	}

	// 根据策略判断
	switch ad.config.Policy {
	case "strict":
		return ad.strictPolicy(cached, alert)
	case "loose":
		return ad.loosePolicy(cached, alert)
	case "custom":
		return ad.customPolicy(cached, alert)
	default:
		return ad.strictPolicy(cached, alert)
	}
}

// 严格策略：只有第一次和状态变化时发送
func (ad *AlertDeduplicator) strictPolicy(cached *CachedAlert, alert *StandardAlert) *DeduplicationResult {
	// 状态变化时发送
	if cached.LastAlert != nil && cached.LastAlert.Status != alert.Status {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "status_changed",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     fmt.Sprintf("告警状态从 %s 变为 %s", cached.LastAlert.Status, alert.Status),
		}
	}

	// 超过最大计数时抑制
	if cached.Count > ad.config.MaxCount {
		// 设置抑制时间
		cached.SuppressUntil = time.Now().Add(ad.config.TimeWindow)
		return &DeduplicationResult{
			ShouldSend: false,
			Action:     "max_count_exceeded",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     fmt.Sprintf("超过最大重复次数 %d，抑制发送", ad.config.MaxCount),
		}
	}

	// 重复告警不发送
	return &DeduplicationResult{
		ShouldSend: false,
		Action:     "duplicate",
		Count:      cached.Count,
		Cached:     cached,
		Reason:     fmt.Sprintf("重复告警，第 %d 次出现", cached.Count),
	}
}

// 宽松策略：允许一定频率的重复发送
func (ad *AlertDeduplicator) loosePolicy(cached *CachedAlert, alert *StandardAlert) *DeduplicationResult {
	// 状态变化时发送
	if cached.LastAlert != nil && cached.LastAlert.Status != alert.Status {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "status_changed",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     "告警状态变化",
		}
	}

	// 每隔一定次数发送一次
	sendInterval := ad.config.MaxCount / 2
	if sendInterval < 1 {
		sendInterval = 1
	}

	if cached.Count%sendInterval == 0 {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "interval_send",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     fmt.Sprintf("间隔发送，第 %d 次", cached.Count),
		}
	}

	return &DeduplicationResult{
		ShouldSend: false,
		Action:     "duplicate",
		Count:      cached.Count,
		Cached:     cached,
		Reason:     "重复告警，等待间隔发送",
	}
}

// 自定义策略：基于告警级别的不同处理
func (ad *AlertDeduplicator) customPolicy(cached *CachedAlert, alert *StandardAlert) *DeduplicationResult {
	// 状态变化时发送
	if cached.LastAlert != nil && cached.LastAlert.Status != alert.Status {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "status_changed",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     "告警状态变化",
		}
	}

	// 根据严重级别调整策略
	var maxCount int
	switch alert.Severity {
	case "critical":
		maxCount = ad.config.MaxCount * 2 // 严重告警允许更多重复
	case "warning":
		maxCount = ad.config.MaxCount
	case "info":
		maxCount = ad.config.MaxCount / 2 // 信息告警减少重复
	default:
		maxCount = ad.config.MaxCount
	}

	if maxCount < 1 {
		maxCount = 1
	}

	if cached.Count <= maxCount {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "severity_based",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     fmt.Sprintf("基于严重级别 %s 的策略发送", alert.Severity),
		}
	}

	return &DeduplicationResult{
		ShouldSend: false,
		Action:     "duplicate",
		Count:      cached.Count,
		Cached:     cached,
		Reason:     fmt.Sprintf("超过严重级别 %s 的最大次数 %d", alert.Severity, maxCount),
	}
}

// 启动清理定时器
func (ad *AlertDeduplicator) startCleaner() {
	cleanupInterval := 5 * time.Minute
	if GlobalConfigManager != nil && GlobalConfigManager.GetCacheConfig() != nil {
		cleanupInterval = GlobalConfigManager.GetCacheConfig().CleanupInterval
	}

	ad.cleaner = time.NewTicker(cleanupInterval)
	go func() {
		for range ad.cleaner.C {
			ad.cleanup()
		}
	}()
}

// 清理过期缓存
func (ad *AlertDeduplicator) cleanup() {
	ad.mutex.Lock()
	defer ad.mutex.Unlock()

	now := time.Now()
	expiredKeys := make([]string, 0)

	for key, cached := range ad.cache {
		// 检查是否过期
		if now.Sub(cached.LastSeen) > ad.config.TimeWindow*2 {
			expiredKeys = append(expiredKeys, key)
		}
	}

	// 删除过期缓存
	for _, key := range expiredKeys {
		delete(ad.cache, key)
	}

	if len(expiredKeys) > 0 {
		logs.Info("[Deduplicator] 清理过期缓存 %d 条", len(expiredKeys))
	}

	// 清理数据库中的过期记录
	go func() {
		err := CleanExpiredDeduplicationRecords(ad.config.TimeWindow * 24) // 保留24个时间窗口的数据
		if err != nil {
			logs.Error("[Deduplicator] 清理数据库过期记录失败: %v", err)
		}
	}()
}

// 持久化到数据库
func (ad *AlertDeduplicator) persistToDatabase(cached *CachedAlert) {
	labelsJSON := cached.LastAlert.GetLabelsString()
	err := AddDeduplicationRecord(
		cached.Fingerprint.Hash,
		cached.LastAlert.AlertName,
		cached.LastAlert.Instance,
		labelsJSON,
	)
	if err != nil {
		logs.Error("[Deduplicator] 持久化去重记录失败: %v", err)
	}
}

// 更新数据库
func (ad *AlertDeduplicator) updateDatabase(cached *CachedAlert) {
	err := UpdateDeduplicationRecord(cached.Fingerprint.Hash, cached.Count)
	if err != nil {
		logs.Error("[Deduplicator] 更新去重记录失败: %v", err)
	}

	// 如果设置了抑制时间，也更新到数据库
	if !cached.SuppressUntil.IsZero() {
		err = SetDeduplicationSuppressUntil(cached.Fingerprint.Hash, cached.SuppressUntil)
		if err != nil {
			logs.Error("[Deduplicator] 更新抑制时间失败: %v", err)
		}
	}
}

// 更新统计信息
func (ad *AlertDeduplicator) updateStats(action string, count int) {
	ad.statsMutex.Lock()
	defer ad.statsMutex.Unlock()

	switch action {
	case "new":
		ad.stats.TotalRecords += count
		ad.stats.ActiveRecords += count
		ad.stats.TodayRecords += count
	case "duplicate":
		ad.stats.TotalDuplicates += count
	}
}

// 获取缓存大小
func (ad *AlertDeduplicator) GetCacheSize() int {
	ad.mutex.RLock()
	defer ad.mutex.RUnlock()
	return len(ad.cache)
}

// 获取统计信息
func (ad *AlertDeduplicator) GetStats() *DeduplicationStats {
	ad.statsMutex.RLock()
	defer ad.statsMutex.RUnlock()

	// 从数据库获取最新统计
	dbStats, err := GetDeduplicationStats()
	if err == nil {
		return dbStats
	}

	// 返回内存统计
	return &DeduplicationStats{
		TotalRecords:    ad.stats.TotalRecords,
		ActiveRecords:   ad.stats.ActiveRecords,
		TodayRecords:    ad.stats.TodayRecords,
		TotalDuplicates: ad.stats.TotalDuplicates,
	}
}

// 获取缓存中的告警
func (ad *AlertDeduplicator) GetCachedAlert(fingerprint string) (*CachedAlert, bool) {
	ad.mutex.RLock()
	defer ad.mutex.RUnlock()
	cached, exists := ad.cache[fingerprint]
	return cached, exists
}

// 手动清除缓存
func (ad *AlertDeduplicator) ClearCache() {
	ad.mutex.Lock()
	defer ad.mutex.Unlock()
	ad.cache = make(map[string]*CachedAlert)
	logs.Info("[Deduplicator] 手动清除所有缓存")
}

// 手动抑制告警
func (ad *AlertDeduplicator) SuppressAlert(fingerprint string, duration time.Duration) error {
	ad.mutex.Lock()
	defer ad.mutex.Unlock()

	cached, exists := ad.cache[fingerprint]
	if !exists {
		return fmt.Errorf("告警指纹不存在: %s", fingerprint)
	}

	cached.SuppressUntil = time.Now().Add(duration)
	
	// 更新数据库
	go ad.updateDatabase(cached)

	logs.Info("[Deduplicator] 手动抑制告警 %s，持续时间 %v", fingerprint, duration)
	return nil
}

// 取消抑制
func (ad *AlertDeduplicator) UnsuppressAlert(fingerprint string) error {
	ad.mutex.Lock()
	defer ad.mutex.Unlock()

	cached, exists := ad.cache[fingerprint]
	if !exists {
		return fmt.Errorf("告警指纹不存在: %s", fingerprint)
	}

	cached.SuppressUntil = time.Time{}
	
	// 更新数据库
	go ad.updateDatabase(cached)

	logs.Info("[Deduplicator] 取消抑制告警 %s", fingerprint)
	return nil
}

// 停止去重管理器
func (ad *AlertDeduplicator) Stop() {
	if ad.cleaner != nil {
		ad.cleaner.Stop()
	}
	logs.Info("[Deduplicator] 去重管理器已停止")
}

// 重新加载配置
func (ad *AlertDeduplicator) ReloadConfig(config *DeduplicationConfig) {
	ad.mutex.Lock()
	defer ad.mutex.Unlock()
	ad.config = config
	logs.Info("[Deduplicator] 配置已重新加载")
}

// 获取配置
func (ad *AlertDeduplicator) GetConfig() *DeduplicationConfig {
	return ad.config
}

// 检查是否启用
func (ad *AlertDeduplicator) IsEnabled() bool {
	return ad.config.Enabled
}

// 获取所有缓存的告警
func (ad *AlertDeduplicator) GetAllCachedAlerts() map[string]*CachedAlert {
	ad.mutex.RLock()
	defer ad.mutex.RUnlock()

	result := make(map[string]*CachedAlert)
	for k, v := range ad.cache {
		result[k] = v
	}
	return result
}

// 导出缓存状态
func (ad *AlertDeduplicator) ExportCacheState() (string, error) {
	ad.mutex.RLock()
	defer ad.mutex.RUnlock()

	data := make(map[string]interface{})
	data["cache_size"] = len(ad.cache)
	data["config"] = ad.config
	data["stats"] = ad.stats
	data["cached_alerts"] = ad.cache

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}