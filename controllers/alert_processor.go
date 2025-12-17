package controllers

import (
	"PrometheusAlert/models"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

// 统一告警处理器
type AlertProcessor struct {
	deduplicator *models.AlertDeduplicator
	aggregator   *models.AlertAggregator
	normalizer   *models.AlertNormalizer
	enabled      bool
}

// 全局告警处理器实例
var GlobalAlertProcessor *AlertProcessor

// 初始化告警处理器
func InitAlertProcessor() {
	config := models.GetGlobalConfig()
	
	GlobalAlertProcessor = &AlertProcessor{
		deduplicator: models.NewAlertDeduplicator(config.Deduplication, config.Fingerprint),
		aggregator:   models.NewAlertAggregator(config.Aggregation),
		normalizer:   models.NewAlertNormalizer(),
		enabled:      config.Deduplication.Enabled,
	}
	
	logs.Info("[AlertProcessor] 告警处理器初始化完成，去重功能: %v", GlobalAlertProcessor.enabled)
}

// 获取全局告警处理器
func GetGlobalAlertProcessor() *AlertProcessor {
	if GlobalAlertProcessor == nil {
		InitAlertProcessor()
	}
	return GlobalAlertProcessor
}

// 处理告警
func (ap *AlertProcessor) ProcessAlert(rawAlert interface{}, source string) (*models.DeduplicationResult, error) {
	if !ap.enabled {
		return &models.DeduplicationResult{
			ShouldSend: true,
			Action:     "disabled",
			Count:      1,
			Reason:     "去重功能已禁用",
		}, nil
	}

	// 1. 标准化告警
	standardAlert, err := ap.normalizer.Normalize(rawAlert, source)
	if err != nil {
		logs.Error("[AlertProcessor] 标准化告警失败: %v", err)
		return &models.DeduplicationResult{
			ShouldSend: true,
			Action:     "normalize_error",
			Count:      1,
			Reason:     fmt.Sprintf("标准化失败: %v", err),
		}, err
	}

	// 2. 去重检查
	result, err := ap.deduplicator.ShouldSend(standardAlert)
	if err != nil {
		logs.Error("[AlertProcessor] 去重检查失败: %v", err)
		return &models.DeduplicationResult{
			ShouldSend: true,
			Action:     "dedup_error",
			Count:      1,
			Reason:     fmt.Sprintf("去重检查失败: %v", err),
		}, err
	}

	// 3. 聚合处理（如果去重检查通过）
	if result.ShouldSend && ap.aggregator.IsEnabled() {
		aggResult, err := ap.aggregator.AddAlert(standardAlert)
		if err != nil {
			logs.Error("[AlertProcessor] 聚合处理失败: %v", err)
		} else if !aggResult.ShouldFlush {
			// 告警被聚合，不立即发送
			logs.Info("[AlertProcessor][%s] 告警已聚合: %s, 组: %s, 数量: %d", 
				source, standardAlert.AlertName, aggResult.Group.GroupKey, aggResult.Group.Count)
			
			return &models.DeduplicationResult{
				ShouldSend: false,
				Action:     "aggregated",
				Count:      result.Count,
				Reason:     aggResult.Reason,
			}, nil
		} else {
			// 聚合组需要刷新，生成聚合告警
			aggregatedAlert := ap.aggregator.GenerateAggregatedMessage(aggResult.Group)
			logs.Info("[AlertProcessor][%s] 聚合组刷新: %s, 包含 %d 个告警", 
				source, aggResult.Group.GroupKey, aggResult.Group.Count)
			
			// 这里可以处理聚合告警的发送
			// 暂时返回原始结果，后续可以扩展为返回聚合告警
			_ = aggregatedAlert
		}
	}

	// 4. 记录处理结果
	ap.logProcessResult(standardAlert, result)

	return result, nil
}

// 记录处理结果
func (ap *AlertProcessor) logProcessResult(alert *models.StandardAlert, result *models.DeduplicationResult) {
	logSign := fmt.Sprintf("[AlertProcessor][%s]", alert.Source)
	
	if result.ShouldSend {
		logs.Info("%s 告警将被发送: %s, 动作: %s, 次数: %d, 原因: %s", 
			logSign, alert.AlertName, result.Action, result.Count, result.Reason)
	} else {
		logs.Debug("%s 告警被抑制: %s, 动作: %s, 次数: %d, 原因: %s", 
			logSign, alert.AlertName, result.Action, result.Count, result.Reason)
	}
}

// 检查是否启用
func (ap *AlertProcessor) IsEnabled() bool {
	return ap.enabled
}

// 启用去重功能
func (ap *AlertProcessor) Enable() {
	ap.enabled = true
	logs.Info("[AlertProcessor] 去重功能已启用")
}

// 禁用去重功能
func (ap *AlertProcessor) Disable() {
	ap.enabled = false
	logs.Info("[AlertProcessor] 去重功能已禁用")
}

// 重新加载配置
func (ap *AlertProcessor) ReloadConfig() error {
	config := models.GetGlobalConfig()
	
	// 重新创建去重管理器
	if ap.deduplicator != nil {
		ap.deduplicator.Stop()
	}
	
	// 重新创建聚合管理器
	if ap.aggregator != nil {
		ap.aggregator.Stop()
	}
	
	ap.deduplicator = models.NewAlertDeduplicator(config.Deduplication, config.Fingerprint)
	ap.aggregator = models.NewAlertAggregator(config.Aggregation)
	ap.enabled = config.Deduplication.Enabled
	
	logs.Info("[AlertProcessor] 配置已重新加载，去重功能: %v, 聚合功能: %v", 
		ap.enabled, ap.aggregator.IsEnabled())
	return nil
}

// 获取统计信息
func (ap *AlertProcessor) GetStats() *models.DeduplicationStats {
	if ap.deduplicator == nil {
		return &models.DeduplicationStats{}
	}
	return ap.deduplicator.GetStats()
}

// 获取缓存大小
func (ap *AlertProcessor) GetCacheSize() int {
	if ap.deduplicator == nil {
		return 0
	}
	return ap.deduplicator.GetCacheSize()
}

// 清除缓存
func (ap *AlertProcessor) ClearCache() {
	if ap.deduplicator != nil {
		ap.deduplicator.ClearCache()
		logs.Info("[AlertProcessor] 缓存已清除")
	}
}

// 手动抑制告警
func (ap *AlertProcessor) SuppressAlert(fingerprint string, duration string) error {
	if ap.deduplicator == nil {
		return fmt.Errorf("去重管理器未初始化")
	}
	
	// 解析持续时间
	d, err := parseDuration(duration)
	if err != nil {
		return fmt.Errorf("无效的持续时间格式: %s", duration)
	}
	
	return ap.deduplicator.SuppressAlert(fingerprint, d)
}

// 取消抑制
func (ap *AlertProcessor) UnsuppressAlert(fingerprint string) error {
	if ap.deduplicator == nil {
		return fmt.Errorf("去重管理器未初始化")
	}
	
	return ap.deduplicator.UnsuppressAlert(fingerprint)
}

// 获取所有缓存的告警
func (ap *AlertProcessor) GetAllCachedAlerts() map[string]*models.CachedAlert {
	if ap.deduplicator == nil {
		return make(map[string]*models.CachedAlert)
	}
	return ap.deduplicator.GetAllCachedAlerts()
}

// 停止告警处理器
func (ap *AlertProcessor) Stop() {
	if ap.deduplicator != nil {
		ap.deduplicator.Stop()
	}
	if ap.aggregator != nil {
		ap.aggregator.Stop()
	}
	logs.Info("[AlertProcessor] 告警处理器已停止")
}

// 获取聚合统计信息
func (ap *AlertProcessor) GetAggregationStats() *models.AggregationStats {
	if ap.aggregator == nil {
		return &models.AggregationStats{}
	}
	return ap.aggregator.GetStats()
}

// 获取活跃聚合组数量
func (ap *AlertProcessor) GetActiveGroupCount() int {
	if ap.aggregator == nil {
		return 0
	}
	return ap.aggregator.GetActiveGroupCount()
}

// 获取所有活跃聚合组
func (ap *AlertProcessor) GetAllActiveGroups() map[string]*models.AggregationGroup {
	if ap.aggregator == nil {
		return make(map[string]*models.AggregationGroup)
	}
	return ap.aggregator.GetAllActiveGroups()
}

// 手动刷新聚合组
func (ap *AlertProcessor) FlushAggregationGroup(groupKey string) (*models.AggregatedAlert, error) {
	if ap.aggregator == nil {
		return nil, fmt.Errorf("聚合管理器未初始化")
	}
	return ap.aggregator.FlushGroup(groupKey)
}

// 清除所有聚合组
func (ap *AlertProcessor) ClearAllAggregationGroups() {
	if ap.aggregator != nil {
		ap.aggregator.ClearAllGroups()
		logs.Info("[AlertProcessor] 所有聚合组已清除")
	}
}

// 启用聚合功能
func (ap *AlertProcessor) EnableAggregation() {
	if ap.aggregator != nil {
		ap.aggregator.Enable()
	}
}

// 禁用聚合功能
func (ap *AlertProcessor) DisableAggregation() {
	if ap.aggregator != nil {
		ap.aggregator.Disable()
	}
}

// 检查聚合功能是否启用
func (ap *AlertProcessor) IsAggregationEnabled() bool {
	if ap.aggregator == nil {
		return false
	}
	return ap.aggregator.IsEnabled()
}

// 解析持续时间字符串
func parseDuration(duration string) (time.Duration, error) {
	// 支持的格式: 5m, 1h, 30s, 2h30m
	return time.ParseDuration(duration)
}

// 辅助函数：处理Prometheus告警
func ProcessPrometheusAlert(rawAlert interface{}) (*models.DeduplicationResult, error) {
	processor := GetGlobalAlertProcessor()
	return processor.ProcessAlert(rawAlert, "prometheus")
}

// 辅助函数：处理阿里云告警
func ProcessAliyunAlert(rawAlert interface{}) (*models.DeduplicationResult, error) {
	processor := GetGlobalAlertProcessor()
	return processor.ProcessAlert(rawAlert, "aliyun")
}

// 辅助函数：处理Zabbix告警
func ProcessZabbixAlert(rawAlert interface{}) (*models.DeduplicationResult, error) {
	processor := GetGlobalAlertProcessor()
	return processor.ProcessAlert(rawAlert, "zabbix")
}