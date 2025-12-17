package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

// 去重配置
type DeduplicationConfig struct {
	Enabled         bool          `json:"enabled"`
	TimeWindow      time.Duration `json:"time_window"`
	MaxCount        int           `json:"max_count"`
	SuppressResolved bool         `json:"suppress_resolved"`
	GroupByLabels   []string      `json:"group_by_labels"`
	Policy          string        `json:"policy"` // strict, loose, custom
}

// 聚合配置
type AggregationConfig struct {
	Enabled       bool          `json:"enabled"`
	TimeWindow    time.Duration `json:"time_window"`
	MaxAlerts     int           `json:"max_alerts"`
	GroupByLabels []string      `json:"group_by_labels"`
	Strategy      string        `json:"strategy"`      // count, list, summary
	FlushInterval time.Duration `json:"flush_interval"`
}

// 缓存配置
type CacheConfig struct {
	Type            string        `json:"type"`             // memory, redis
	MaxSize         int           `json:"max_size"`
	TTL             time.Duration `json:"ttl"`
	CleanupInterval time.Duration `json:"cleanup_interval"`
	RedisAddr       string        `json:"redis_addr"`
	RedisPassword   string        `json:"redis_password"`
	RedisDB         int           `json:"redis_db"`
}

// 去重聚合总配置
type DeduplicationAggregationConfig struct {
	Deduplication *DeduplicationConfig `json:"deduplication"`
	Aggregation   *AggregationConfig   `json:"aggregation"`
	Cache         *CacheConfig         `json:"cache"`
	Fingerprint   *FingerprintConfig   `json:"fingerprint"`
}

// 配置管理器
type ConfigManager struct {
	config   *DeduplicationAggregationConfig
	filePath string
}

// 创建配置管理器
func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		config:   getDefaultConfig(),
		filePath: "conf/deduplication.conf",
	}
}

// 获取默认配置
func getDefaultConfig() *DeduplicationAggregationConfig {
	return &DeduplicationAggregationConfig{
		Deduplication: &DeduplicationConfig{
			Enabled:         true,
			TimeWindow:      5 * time.Minute,
			MaxCount:        5,
			SuppressResolved: true,
			GroupByLabels:   []string{"alertname", "instance", "severity"},
			Policy:          "strict",
		},
		Aggregation: &AggregationConfig{
			Enabled:       false, // 默认关闭聚合
			TimeWindow:    1 * time.Minute,
			MaxAlerts:     10,
			GroupByLabels: []string{"alertname", "severity"},
			Strategy:      "summary",
			FlushInterval: 30 * time.Second,
		},
		Cache: &CacheConfig{
			Type:            "memory",
			MaxSize:         10000,
			TTL:             1 * time.Hour,
			CleanupInterval: 5 * time.Minute,
			RedisAddr:       "localhost:6379",
			RedisPassword:   "",
			RedisDB:         0,
		},
		Fingerprint: &FingerprintConfig{
			Algorithm:     "md5",
			IncludeFields: []string{"alert_name", "instance", "labels"},
			ExcludeLabels: []string{"__name__", "__tmp_", "receive_time"},
			IncludeLabels: []string{},
		},
	}
}

// 从Beego配置加载
func (cm *ConfigManager) LoadFromBeegoConfig() error {
	// 去重配置
	if enabled := beego.AppConfig.String("deduplication::enabled"); enabled != "" {
		cm.config.Deduplication.Enabled = enabled == "true"
	}
	
	if timeWindow := beego.AppConfig.String("deduplication::time_window"); timeWindow != "" {
		if duration, err := time.ParseDuration(timeWindow); err == nil {
			cm.config.Deduplication.TimeWindow = duration
		}
	}
	
	if maxCount, err := beego.AppConfig.Int("deduplication::max_count"); err == nil {
		cm.config.Deduplication.MaxCount = maxCount
	}
	
	if suppressResolved := beego.AppConfig.String("deduplication::suppress_resolved"); suppressResolved != "" {
		cm.config.Deduplication.SuppressResolved = suppressResolved == "true"
	}
	
	if groupByLabels := beego.AppConfig.String("deduplication::group_by_labels"); groupByLabels != "" {
		cm.config.Deduplication.GroupByLabels = parseStringSlice(groupByLabels)
	}
	
	if policy := beego.AppConfig.String("deduplication::policy"); policy != "" {
		cm.config.Deduplication.Policy = policy
	}
	
	// 聚合配置
	if enabled := beego.AppConfig.String("aggregation::enabled"); enabled != "" {
		cm.config.Aggregation.Enabled = enabled == "true"
	}
	
	if timeWindow := beego.AppConfig.String("aggregation::time_window"); timeWindow != "" {
		if duration, err := time.ParseDuration(timeWindow); err == nil {
			cm.config.Aggregation.TimeWindow = duration
		}
	}
	
	if maxAlerts, err := beego.AppConfig.Int("aggregation::max_alerts"); err == nil {
		cm.config.Aggregation.MaxAlerts = maxAlerts
	}
	
	if groupByLabels := beego.AppConfig.String("aggregation::group_by_labels"); groupByLabels != "" {
		cm.config.Aggregation.GroupByLabels = parseStringSlice(groupByLabels)
	}
	
	if strategy := beego.AppConfig.String("aggregation::strategy"); strategy != "" {
		cm.config.Aggregation.Strategy = strategy
	}
	
	if flushInterval := beego.AppConfig.String("aggregation::flush_interval"); flushInterval != "" {
		if duration, err := time.ParseDuration(flushInterval); err == nil {
			cm.config.Aggregation.FlushInterval = duration
		}
	}
	
	// 缓存配置
	if cacheType := beego.AppConfig.String("cache::type"); cacheType != "" {
		cm.config.Cache.Type = cacheType
	}
	
	if maxSize, err := beego.AppConfig.Int("cache::max_size"); err == nil {
		cm.config.Cache.MaxSize = maxSize
	}
	
	if ttl := beego.AppConfig.String("cache::ttl"); ttl != "" {
		if duration, err := time.ParseDuration(ttl); err == nil {
			cm.config.Cache.TTL = duration
		}
	}
	
	if cleanupInterval := beego.AppConfig.String("cache::cleanup_interval"); cleanupInterval != "" {
		if duration, err := time.ParseDuration(cleanupInterval); err == nil {
			cm.config.Cache.CleanupInterval = duration
		}
	}
	
	if redisAddr := beego.AppConfig.String("cache::redis_addr"); redisAddr != "" {
		cm.config.Cache.RedisAddr = redisAddr
	}
	
	if redisPassword := beego.AppConfig.String("cache::redis_password"); redisPassword != "" {
		cm.config.Cache.RedisPassword = redisPassword
	}
	
	if redisDB, err := beego.AppConfig.Int("cache::redis_db"); err == nil {
		cm.config.Cache.RedisDB = redisDB
	}
	
	// 指纹配置
	if algorithm := beego.AppConfig.String("fingerprint::algorithm"); algorithm != "" {
		cm.config.Fingerprint.Algorithm = algorithm
	}
	
	if includeFields := beego.AppConfig.String("fingerprint::include_fields"); includeFields != "" {
		cm.config.Fingerprint.IncludeFields = parseStringSlice(includeFields)
	}
	
	if excludeLabels := beego.AppConfig.String("fingerprint::exclude_labels"); excludeLabels != "" {
		cm.config.Fingerprint.ExcludeLabels = parseStringSlice(excludeLabels)
	}
	
	if includeLabels := beego.AppConfig.String("fingerprint::include_labels"); includeLabels != "" {
		cm.config.Fingerprint.IncludeLabels = parseStringSlice(includeLabels)
	}
	
	return nil
}

// 解析字符串切片
func parseStringSlice(str string) []string {
	if str == "" {
		return []string{}
	}
	
	var result []string
	err := json.Unmarshal([]byte(fmt.Sprintf(`["%s"]`, str)), &result)
	if err != nil {
		// 如果JSON解析失败，尝试逗号分割
		result = []string{}
		for _, item := range []string{str} {
			if item != "" {
				result = append(result, item)
			}
		}
	}
	return result
}

// 获取配置
func (cm *ConfigManager) GetConfig() *DeduplicationAggregationConfig {
	return cm.config
}

// 获取去重配置
func (cm *ConfigManager) GetDeduplicationConfig() *DeduplicationConfig {
	return cm.config.Deduplication
}

// 获取聚合配置
func (cm *ConfigManager) GetAggregationConfig() *AggregationConfig {
	return cm.config.Aggregation
}

// 获取缓存配置
func (cm *ConfigManager) GetCacheConfig() *CacheConfig {
	return cm.config.Cache
}

// 获取指纹配置
func (cm *ConfigManager) GetFingerprintConfig() *FingerprintConfig {
	return cm.config.Fingerprint
}

// 验证配置
func (cm *ConfigManager) ValidateConfig() error {
	config := cm.config
	
	// 验证去重配置
	if config.Deduplication.TimeWindow <= 0 {
		return fmt.Errorf("deduplication time_window must be positive")
	}
	
	if config.Deduplication.MaxCount <= 0 {
		return fmt.Errorf("deduplication max_count must be positive")
	}
	
	if config.Deduplication.Policy != "strict" && config.Deduplication.Policy != "loose" && config.Deduplication.Policy != "custom" {
		return fmt.Errorf("deduplication policy must be one of: strict, loose, custom")
	}
	
	// 验证聚合配置
	if config.Aggregation.Enabled {
		if config.Aggregation.TimeWindow <= 0 {
			return fmt.Errorf("aggregation time_window must be positive")
		}
		
		if config.Aggregation.MaxAlerts <= 0 {
			return fmt.Errorf("aggregation max_alerts must be positive")
		}
		
		if config.Aggregation.Strategy != "count" && config.Aggregation.Strategy != "list" && config.Aggregation.Strategy != "summary" {
			return fmt.Errorf("aggregation strategy must be one of: count, list, summary")
		}
		
		if config.Aggregation.FlushInterval <= 0 {
			return fmt.Errorf("aggregation flush_interval must be positive")
		}
	}
	
	// 验证缓存配置
	if config.Cache.Type != "memory" && config.Cache.Type != "redis" {
		return fmt.Errorf("cache type must be one of: memory, redis")
	}
	
	if config.Cache.MaxSize <= 0 {
		return fmt.Errorf("cache max_size must be positive")
	}
	
	if config.Cache.TTL <= 0 {
		return fmt.Errorf("cache ttl must be positive")
	}
	
	// 验证指纹配置
	if config.Fingerprint.Algorithm != "md5" && config.Fingerprint.Algorithm != "sha256" {
		return fmt.Errorf("fingerprint algorithm must be one of: md5, sha256")
	}
	
	if len(config.Fingerprint.IncludeFields) == 0 {
		return fmt.Errorf("fingerprint include_fields cannot be empty")
	}
	
	return nil
}

// 转换为JSON
func (cm *ConfigManager) ToJSON() (string, error) {
	data, err := json.MarshalIndent(cm.config, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// 从JSON加载
func (cm *ConfigManager) FromJSON(jsonStr string) error {
	config := &DeduplicationAggregationConfig{}
	err := json.Unmarshal([]byte(jsonStr), config)
	if err != nil {
		return err
	}
	
	cm.config = config
	return cm.ValidateConfig()
}

// 重置为默认配置
func (cm *ConfigManager) ResetToDefault() {
	cm.config = getDefaultConfig()
}

// 更新去重配置
func (cm *ConfigManager) UpdateDeduplicationConfig(config *DeduplicationConfig) error {
	cm.config.Deduplication = config
	return cm.ValidateConfig()
}

// 更新聚合配置
func (cm *ConfigManager) UpdateAggregationConfig(config *AggregationConfig) error {
	cm.config.Aggregation = config
	return cm.ValidateConfig()
}

// 更新缓存配置
func (cm *ConfigManager) UpdateCacheConfig(config *CacheConfig) error {
	cm.config.Cache = config
	return cm.ValidateConfig()
}

// 更新指纹配置
func (cm *ConfigManager) UpdateFingerprintConfig(config *FingerprintConfig) error {
	cm.config.Fingerprint = config
	return cm.ValidateConfig()
}

// 获取配置摘要
func (cm *ConfigManager) GetConfigSummary() *ConfigSummary {
	return &ConfigSummary{
		DeduplicationEnabled: cm.config.Deduplication.Enabled,
		AggregationEnabled:   cm.config.Aggregation.Enabled,
		CacheType:           cm.config.Cache.Type,
		FingerprintAlgorithm: cm.config.Fingerprint.Algorithm,
		TimeWindow:          cm.config.Deduplication.TimeWindow,
		MaxCount:            cm.config.Deduplication.MaxCount,
	}
}

// 配置摘要
type ConfigSummary struct {
	DeduplicationEnabled bool          `json:"deduplication_enabled"`
	AggregationEnabled   bool          `json:"aggregation_enabled"`
	CacheType           string        `json:"cache_type"`
	FingerprintAlgorithm string        `json:"fingerprint_algorithm"`
	TimeWindow          time.Duration `json:"time_window"`
	MaxCount            int           `json:"max_count"`
}

// 全局配置管理器实例
var GlobalConfigManager *ConfigManager

// 初始化全局配置管理器
func InitConfigManager() error {
	GlobalConfigManager = NewConfigManager()
	return GlobalConfigManager.LoadFromBeegoConfig()
}

// 获取全局配置
func GetGlobalConfig() *DeduplicationAggregationConfig {
	if GlobalConfigManager == nil {
		InitConfigManager()
	}
	return GlobalConfigManager.GetConfig()
}