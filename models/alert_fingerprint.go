package models

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"time"
)

// 告警指纹
type AlertFingerprint struct {
	Hash      string            `json:"hash"`
	Labels    map[string]string `json:"labels"`
	AlertName string            `json:"alert_name"`
	Instance  string            `json:"instance"`
	CreatedAt time.Time         `json:"created_at"`
}

// 指纹配置
type FingerprintConfig struct {
	Algorithm     string   `json:"algorithm"`      // md5, sha256
	IncludeFields []string `json:"include_fields"` // 参与指纹计算的字段
	ExcludeLabels []string `json:"exclude_labels"` // 排除的标签
	IncludeLabels []string `json:"include_labels"` // 包含的标签(为空则包含所有)
}

// 告警指纹生成器
type AlertFingerprinter struct {
	config *FingerprintConfig
}

// 创建指纹生成器
func NewAlertFingerprinter(config *FingerprintConfig) *AlertFingerprinter {
	if config == nil {
		config = &FingerprintConfig{
			Algorithm:     "md5",
			IncludeFields: []string{"alert_name", "instance", "labels"},
			ExcludeLabels: []string{"__name__", "__tmp_", "receive_time"},
			IncludeLabels: []string{},
		}
	}
	
	return &AlertFingerprinter{
		config: config,
	}
}

// 生成告警指纹
func (af *AlertFingerprinter) GenerateFingerprint(alert *StandardAlert) *AlertFingerprint {
	// 构建指纹字符串
	fingerprintStr := af.buildFingerprintString(alert)
	
	// 计算哈希值
	hash := af.calculateHash(fingerprintStr)
	
	return &AlertFingerprint{
		Hash:      hash,
		Labels:    af.filterLabels(alert.Labels),
		AlertName: alert.AlertName,
		Instance:  alert.Instance,
		CreatedAt: time.Now(),
	}
}

// 构建指纹字符串
func (af *AlertFingerprinter) buildFingerprintString(alert *StandardAlert) string {
	var parts []string
	
	for _, field := range af.config.IncludeFields {
		switch field {
		case "alert_name":
			if alert.AlertName != "" {
				parts = append(parts, fmt.Sprintf("alertname=%s", alert.AlertName))
			}
		case "instance":
			if alert.Instance != "" {
				parts = append(parts, fmt.Sprintf("instance=%s", alert.Instance))
			}
		case "labels":
			labelsStr := af.labelsToString(alert.Labels)
			if labelsStr != "" {
				parts = append(parts, fmt.Sprintf("labels=%s", labelsStr))
			}
		case "severity":
			if alert.Severity != "" {
				parts = append(parts, fmt.Sprintf("severity=%s", alert.Severity))
			}
		case "source":
			if alert.Source != "" {
				parts = append(parts, fmt.Sprintf("source=%s", alert.Source))
			}
		}
	}
	
	return strings.Join(parts, "|")
}

// 标签转字符串
func (af *AlertFingerprinter) labelsToString(labels map[string]string) string {
	if len(labels) == 0 {
		return ""
	}
	
	// 过滤标签
	filteredLabels := af.filterLabels(labels)
	if len(filteredLabels) == 0 {
		return ""
	}
	
	// 排序标签键以确保一致性
	keys := make([]string, 0, len(filteredLabels))
	for key := range filteredLabels {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	
	// 构建标签字符串
	var labelParts []string
	for _, key := range keys {
		labelParts = append(labelParts, fmt.Sprintf("%s=%s", key, filteredLabels[key]))
	}
	
	return strings.Join(labelParts, ",")
}

// 过滤标签
func (af *AlertFingerprinter) filterLabels(labels map[string]string) map[string]string {
	filtered := make(map[string]string)
	
	for key, value := range labels {
		// 检查是否在排除列表中
		if af.isExcludedLabel(key) {
			continue
		}
		
		// 检查是否在包含列表中(如果包含列表不为空)
		if len(af.config.IncludeLabels) > 0 && !af.isIncludedLabel(key) {
			continue
		}
		
		filtered[key] = value
	}
	
	return filtered
}

// 检查标签是否被排除
func (af *AlertFingerprinter) isExcludedLabel(label string) bool {
	for _, excluded := range af.config.ExcludeLabels {
		if strings.HasPrefix(label, excluded) {
			return true
		}
	}
	return false
}

// 检查标签是否被包含
func (af *AlertFingerprinter) isIncludedLabel(label string) bool {
	for _, included := range af.config.IncludeLabels {
		if label == included || strings.HasPrefix(label, included) {
			return true
		}
	}
	return false
}

// 计算哈希值
func (af *AlertFingerprinter) calculateHash(input string) string {
	switch af.config.Algorithm {
	case "sha256":
		hash := sha256.Sum256([]byte(input))
		return hex.EncodeToString(hash[:])
	case "md5":
		fallthrough
	default:
		hash := md5.Sum([]byte(input))
		return hex.EncodeToString(hash[:])
	}
}

// 验证指纹
func (af *AlertFingerprinter) ValidateFingerprint(fingerprint string) bool {
	switch af.config.Algorithm {
	case "sha256":
		return len(fingerprint) == 64
	case "md5":
		return len(fingerprint) == 32
	default:
		return len(fingerprint) == 32
	}
}

// 比较两个告警是否相同
func (af *AlertFingerprinter) IsSameAlert(alert1, alert2 *StandardAlert) bool {
	fp1 := af.GenerateFingerprint(alert1)
	fp2 := af.GenerateFingerprint(alert2)
	return fp1.Hash == fp2.Hash
}

// 获取指纹统计信息
func (af *AlertFingerprinter) GetFingerprintStats(fingerprints []*AlertFingerprint) *FingerprintStats {
	stats := &FingerprintStats{
		Total:      len(fingerprints),
		Unique:     0,
		Duplicates: 0,
		Algorithm:  af.config.Algorithm,
	}
	
	hashCount := make(map[string]int)
	for _, fp := range fingerprints {
		hashCount[fp.Hash]++
	}
	
	stats.Unique = len(hashCount)
	for _, count := range hashCount {
		if count > 1 {
			stats.Duplicates += count - 1
		}
	}
	
	return stats
}

// 指纹统计信息
type FingerprintStats struct {
	Total      int    `json:"total"`
	Unique     int    `json:"unique"`
	Duplicates int    `json:"duplicates"`
	Algorithm  string `json:"algorithm"`
}

// 批量生成指纹
func (af *AlertFingerprinter) GenerateFingerprints(alerts []*StandardAlert) []*AlertFingerprint {
	fingerprints := make([]*AlertFingerprint, len(alerts))
	for i, alert := range alerts {
		fingerprints[i] = af.GenerateFingerprint(alert)
	}
	return fingerprints
}

// 指纹缓存项
type FingerprintCacheItem struct {
	Fingerprint *AlertFingerprint `json:"fingerprint"`
	CreatedAt   time.Time         `json:"created_at"`
	AccessCount int               `json:"access_count"`
	LastAccess  time.Time         `json:"last_access"`
}

// 指纹缓存
type FingerprintCache struct {
	cache   map[string]*FingerprintCacheItem
	maxSize int
	ttl     time.Duration
}

// 创建指纹缓存
func NewFingerprintCache(maxSize int, ttl time.Duration) *FingerprintCache {
	return &FingerprintCache{
		cache:   make(map[string]*FingerprintCacheItem),
		maxSize: maxSize,
		ttl:     ttl,
	}
}

// 获取缓存的指纹
func (fc *FingerprintCache) Get(key string) (*AlertFingerprint, bool) {
	item, exists := fc.cache[key]
	if !exists {
		return nil, false
	}
	
	// 检查是否过期
	if time.Since(item.CreatedAt) > fc.ttl {
		delete(fc.cache, key)
		return nil, false
	}
	
	// 更新访问信息
	item.AccessCount++
	item.LastAccess = time.Now()
	
	return item.Fingerprint, true
}

// 设置缓存
func (fc *FingerprintCache) Set(key string, fingerprint *AlertFingerprint) {
	// 检查缓存大小
	if len(fc.cache) >= fc.maxSize {
		fc.evictLRU()
	}
	
	fc.cache[key] = &FingerprintCacheItem{
		Fingerprint: fingerprint,
		CreatedAt:   time.Now(),
		AccessCount: 1,
		LastAccess:  time.Now(),
	}
}

// LRU淘汰
func (fc *FingerprintCache) evictLRU() {
	var oldestKey string
	var oldestTime time.Time = time.Now()
	
	for key, item := range fc.cache {
		if item.LastAccess.Before(oldestTime) {
			oldestTime = item.LastAccess
			oldestKey = key
		}
	}
	
	if oldestKey != "" {
		delete(fc.cache, oldestKey)
	}
}

// 清理过期缓存
func (fc *FingerprintCache) Cleanup() {
	now := time.Now()
	for key, item := range fc.cache {
		if now.Sub(item.CreatedAt) > fc.ttl {
			delete(fc.cache, key)
		}
	}
}

// 获取缓存统计
func (fc *FingerprintCache) GetStats() *CacheStats {
	return &CacheStats{
		Size:    len(fc.cache),
		MaxSize: fc.maxSize,
		TTL:     fc.ttl,
	}
}

// 缓存统计
type CacheStats struct {
	Size    int           `json:"size"`
	MaxSize int           `json:"max_size"`
	TTL     time.Duration `json:"ttl"`
}