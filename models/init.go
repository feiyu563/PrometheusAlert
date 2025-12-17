package models

import (
	"github.com/astaxie/beego/orm"
)

// 初始化去重聚合模块
func InitDeduplicationAggregation() error {
	// 注册新的数据模型
	orm.RegisterModel(new(AlertDeduplicationRecord))
	orm.RegisterModel(new(AlertAggregationRecord))
	
	// 初始化配置管理器
	err := InitConfigManager()
	if err != nil {
		return err
	}
	
	// 初始化策略管理器
	InitPolicyManager()
	
	return nil
}

// 获取默认去重管理器
func GetDefaultDeduplicator() *AlertDeduplicator {
	config := GetGlobalConfig()
	return NewAlertDeduplicator(config.Deduplication, config.Fingerprint)
}

// 获取默认内存缓存
func GetDefaultMemoryCache() *MemoryCache {
	config := GetGlobalConfig()
	return NewMemoryCache(config.Cache.MaxSize, config.Cache.TTL)
}