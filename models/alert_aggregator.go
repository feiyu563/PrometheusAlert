package models

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
)

// 聚合组
type AggregationGroup struct {
	GroupKey    string           `json:"group_key"`
	Alerts      []*StandardAlert `json:"alerts"`
	FirstSeen   time.Time        `json:"first_seen"`
	LastSeen    time.Time        `json:"last_seen"`
	Count       int              `json:"count"`
	Status      string           `json:"status"`
	Severity    string           `json:"severity"`
	Labels      map[string]string `json:"labels"`
}

// 聚合结果
type AggregationResult struct {
	ShouldFlush bool              `json:"should_flush"`
	Group       *AggregationGroup `json:"group"`
	Action      string            `json:"action"`
	Reason      string            `json:"reason"`
}

// 告警聚合管理器
type AlertAggregator struct {
	config     *AggregationConfig
	groups     map[string]*AggregationGroup // 分组键 -> 聚合组
	mutex      sync.RWMutex
	flushTimer *time.Ticker // 刷新定时器
	stats      *AggregationStats
	statsMutex sync.RWMutex
}

// 聚合统计信息
type AggregationStats struct {
	TotalGroups     int `json:"total_groups"`
	ActiveGroups    int `json:"active_groups"`
	FlushedGroups   int `json:"flushed_groups"`
	TotalAlerts     int `json:"total_alerts"`
	AverageGroupSize float64 `json:"average_group_size"`
}

// 创建告警聚合管理器
func NewAlertAggregator(config *AggregationConfig) *AlertAggregator {
	if config == nil {
		config = &AggregationConfig{
			Enabled:       false,
			TimeWindow:    1 * time.Minute,
			MaxAlerts:     10,
			GroupByLabels: []string{"alertname", "severity"},
			Strategy:      "summary",
			FlushInterval: 30 * time.Second,
		}
	}

	aggregator := &AlertAggregator{
		config: config,
		groups: make(map[string]*AggregationGroup),
		stats: &AggregationStats{
			TotalGroups:     0,
			ActiveGroups:    0,
			FlushedGroups:   0,
			TotalAlerts:     0,
			AverageGroupSize: 0,
		},
	}

	// 启动刷新定时器
	if config.Enabled {
		aggregator.startFlushTimer()
	}

	return aggregator
}

// 添加告警到聚合组
func (aa *AlertAggregator) AddAlert(alert *StandardAlert) (*AggregationResult, error) {
	if !aa.config.Enabled {
		return &AggregationResult{
			ShouldFlush: true,
			Group:       nil,
			Action:      "disabled",
			Reason:      "聚合功能已禁用",
		}, nil
	}

	groupKey := aa.generateGroupKey(alert)
	
	aa.mutex.Lock()
	defer aa.mutex.Unlock()

	group, exists := aa.groups[groupKey]
	if !exists {
		group = &AggregationGroup{
			GroupKey:  groupKey,
			Alerts:    make([]*StandardAlert, 0),
			FirstSeen: time.Now(),
			Status:    "active",
			Severity:  alert.Severity,
			Labels:    aa.extractGroupLabels(alert),
		}
		aa.groups[groupKey] = group
		aa.updateStats("new_group", 1)
	}

	group.Alerts = append(group.Alerts, alert)
	group.LastSeen = time.Now()
	group.Count = len(group.Alerts)
	
	// 更新组的严重级别（取最高级别）
	group.Severity = aa.getHighestSeverity(group.Alerts)
	
	aa.updateStats("add_alert", 1)

	// 检查是否需要刷新
	shouldFlush := aa.shouldFlushGroup(group)
	
	if shouldFlush {
		// 异步持久化到数据库
		go aa.persistToDatabase(group)
	}

	return &AggregationResult{
		ShouldFlush: shouldFlush,
		Group:       group,
		Action:      "aggregated",
		Reason:      fmt.Sprintf("告警已添加到聚合组 %s，当前数量: %d", groupKey, group.Count),
	}, nil
}

// 生成分组键
func (aa *AlertAggregator) generateGroupKey(alert *StandardAlert) string {
	var keyParts []string
	
	for _, label := range aa.config.GroupByLabels {
		var value string
		switch label {
		case "alertname":
			value = alert.AlertName
		case "severity":
			value = alert.Severity
		case "instance":
			value = alert.Instance
		case "status":
			value = alert.Status
		case "source":
			value = alert.Source
		default:
			// 检查是否是标签字段
			if strings.HasPrefix(label, "labels.") {
				labelKey := strings.TrimPrefix(label, "labels.")
				if labelValue, exists := alert.Labels[labelKey]; exists {
					value = labelValue
				}
			} else if labelValue, exists := alert.Labels[label]; exists {
				value = labelValue
			}
		}
		
		if value != "" {
			keyParts = append(keyParts, fmt.Sprintf("%s=%s", label, value))
		}
	}
	
	// 排序以确保一致性
	sort.Strings(keyParts)
	
	return strings.Join(keyParts, "|")
}

// 提取分组标签
func (aa *AlertAggregator) extractGroupLabels(alert *StandardAlert) map[string]string {
	labels := make(map[string]string)
	
	for _, label := range aa.config.GroupByLabels {
		switch label {
		case "alertname":
			labels["alertname"] = alert.AlertName
		case "severity":
			labels["severity"] = alert.Severity
		case "instance":
			labels["instance"] = alert.Instance
		case "status":
			labels["status"] = alert.Status
		case "source":
			labels["source"] = alert.Source
		default:
			if labelValue, exists := alert.Labels[label]; exists {
				labels[label] = labelValue
			}
		}
	}
	
	return labels
}

// 判断是否应该刷新组
func (aa *AlertAggregator) shouldFlushGroup(group *AggregationGroup) bool {
	// 检查告警数量
	if group.Count >= aa.config.MaxAlerts {
		logs.Debug("[Aggregator] 组 %s 达到最大告警数 %d，触发刷新", group.GroupKey, aa.config.MaxAlerts)
		return true
	}
	
	// 检查时间窗口
	if time.Since(group.FirstSeen) >= aa.config.TimeWindow {
		logs.Debug("[Aggregator] 组 %s 达到时间窗口 %v，触发刷新", group.GroupKey, aa.config.TimeWindow)
		return true
	}
	
	return false
}

// 获取最高严重级别
func (aa *AlertAggregator) getHighestSeverity(alerts []*StandardAlert) string {
	severityOrder := map[string]int{
		"critical": 4,
		"warning":  3,
		"info":     2,
		"":         1,
	}
	
	highestLevel := ""
	highestValue := 0
	
	for _, alert := range alerts {
		if value, exists := severityOrder[alert.Severity]; exists {
			if value > highestValue {
				highestValue = value
				highestLevel = alert.Severity
			}
		}
	}
	
	return highestLevel
}

// 生成聚合消息
func (aa *AlertAggregator) GenerateAggregatedMessage(group *AggregationGroup) *AggregatedAlert {
	aggregated := &AggregatedAlert{
		GroupKey:     group.GroupKey,
		Count:        group.Count,
		FirstSeen:    group.FirstSeen,
		LastSeen:     group.LastSeen,
		Alerts:       group.Alerts,
		IsAggregated: true,
		Severity:     group.Severity,
		Status:       group.Status,
	}
	
	// 根据策略生成摘要和描述
	switch aa.config.Strategy {
	case "count":
		aggregated.Summary = aa.generateCountSummary(group)
		aggregated.Description = aa.generateCountDescription(group)
	case "list":
		aggregated.Summary = aa.generateListSummary(group)
		aggregated.Description = aa.generateListDescription(group)
	case "summary":
		fallthrough
	default:
		aggregated.Summary = aa.generateSummary(group)
		aggregated.Description = aa.generateDescription(group)
	}
	
	return aggregated
}

// 生成计数摘要
func (aa *AlertAggregator) generateCountSummary(group *AggregationGroup) string {
	return fmt.Sprintf("聚合告警: %d 个告警 (%s)", group.Count, group.Severity)
}

// 生成计数描述
func (aa *AlertAggregator) generateCountDescription(group *AggregationGroup) string {
	alertNames := make(map[string]int)
	instances := make(map[string]int)
	
	for _, alert := range group.Alerts {
		alertNames[alert.AlertName]++
		if alert.Instance != "" {
			instances[alert.Instance]++
		}
	}
	
	description := fmt.Sprintf("时间范围: %s - %s\n", 
		group.FirstSeen.Format("2006-01-02 15:04:05"),
		group.LastSeen.Format("2006-01-02 15:04:05"))
	
	description += fmt.Sprintf("告警类型: %d 种\n", len(alertNames))
	description += fmt.Sprintf("影响实例: %d 个\n", len(instances))
	
	return description
}

// 生成列表摘要
func (aa *AlertAggregator) generateListSummary(group *AggregationGroup) string {
	alertNames := make(map[string]int)
	for _, alert := range group.Alerts {
		alertNames[alert.AlertName]++
	}
	
	var names []string
	for name := range alertNames {
		names = append(names, name)
	}
	
	if len(names) > 3 {
		return fmt.Sprintf("聚合告警: %s 等 %d 个告警", strings.Join(names[:3], ", "), len(names))
	}
	
	return fmt.Sprintf("聚合告警: %s", strings.Join(names, ", "))
}

// 生成列表描述
func (aa *AlertAggregator) generateListDescription(group *AggregationGroup) string {
	var descriptions []string
	
	for i, alert := range group.Alerts {
		if i >= 10 { // 最多显示10个
			descriptions = append(descriptions, fmt.Sprintf("... 还有 %d 个告警", group.Count-10))
			break
		}
		
		desc := fmt.Sprintf("%d. %s", i+1, alert.AlertName)
		if alert.Instance != "" {
			desc += fmt.Sprintf(" (%s)", alert.Instance)
		}
		if alert.Summary != "" {
			desc += fmt.Sprintf(": %s", alert.Summary)
		}
		
		descriptions = append(descriptions, desc)
	}
	
	return strings.Join(descriptions, "\n")
}

// 生成摘要
func (aa *AlertAggregator) generateSummary(group *AggregationGroup) string {
	alertNames := make(map[string]int)
	for _, alert := range group.Alerts {
		alertNames[alert.AlertName]++
	}
	
	if len(alertNames) == 1 {
		for name, count := range alertNames {
			if count == 1 {
				return fmt.Sprintf("告警: %s", name)
			}
			return fmt.Sprintf("告警: %s (x%d)", name, count)
		}
	}
	
	return fmt.Sprintf("聚合告警: %d 个告警类型，共 %d 个告警", len(alertNames), group.Count)
}

// 生成描述
func (aa *AlertAggregator) generateDescription(group *AggregationGroup) string {
	alertNames := make(map[string]int)
	instances := make(map[string]int)
	
	for _, alert := range group.Alerts {
		alertNames[alert.AlertName]++
		if alert.Instance != "" {
			instances[alert.Instance]++
		}
	}
	
	description := fmt.Sprintf("聚合时间: %s - %s\n", 
		group.FirstSeen.Format("2006-01-02 15:04:05"),
		group.LastSeen.Format("2006-01-02 15:04:05"))
	
	description += fmt.Sprintf("严重级别: %s\n", group.Severity)
	description += fmt.Sprintf("告警总数: %d\n", group.Count)
	
	// 告警类型统计
	description += "告警类型:\n"
	for name, count := range alertNames {
		description += fmt.Sprintf("  - %s: %d 次\n", name, count)
	}
	
	// 实例统计
	if len(instances) > 0 {
		description += "影响实例:\n"
		count := 0
		for instance, alertCount := range instances {
			if count >= 5 { // 最多显示5个实例
				description += fmt.Sprintf("  ... 还有 %d 个实例\n", len(instances)-5)
				break
			}
			description += fmt.Sprintf("  - %s: %d 个告警\n", instance, alertCount)
			count++
		}
	}
	
	return description
}

// 启动刷新定时器
func (aa *AlertAggregator) startFlushTimer() {
	aa.flushTimer = time.NewTicker(aa.config.FlushInterval)
	go func() {
		for range aa.flushTimer.C {
			aa.flushExpiredGroups()
		}
	}()
}

// 刷新过期组
func (aa *AlertAggregator) flushExpiredGroups() {
	aa.mutex.Lock()
	defer aa.mutex.Unlock()

	now := time.Now()
	expiredGroups := make([]*AggregationGroup, 0)
	
	for key, group := range aa.groups {
		if now.Sub(group.FirstSeen) >= aa.config.TimeWindow {
			expiredGroups = append(expiredGroups, group)
			delete(aa.groups, key)
		}
	}
	
	if len(expiredGroups) > 0 {
		logs.Info("[Aggregator] 刷新过期聚合组 %d 个", len(expiredGroups))
		aa.updateStats("flush_expired", len(expiredGroups))
		
		// 异步处理过期组
		go func() {
			for _, group := range expiredGroups {
				aa.persistToDatabase(group)
			}
		}()
	}
}

// 持久化到数据库
func (aa *AlertAggregator) persistToDatabase(group *AggregationGroup) {
	// 序列化告警数据
	alertsData, err := json.Marshal(group.Alerts)
	if err != nil {
		logs.Error("[Aggregator] 序列化告警数据失败: %v", err)
		return
	}
	
	// 序列化分组标签
	labelsData, err := json.Marshal(group.Labels)
	if err != nil {
		logs.Error("[Aggregator] 序列化分组标签失败: %v", err)
		return
	}
	
	err = AddAggregationRecord(
		group.GroupKey,
		string(labelsData),
		group.Count,
		group.FirstSeen,
		group.LastSeen,
		group.Status,
		aa.generateSummary(group),
		aa.generateDescription(group),
		string(alertsData),
	)
	
	if err != nil {
		logs.Error("[Aggregator] 持久化聚合记录失败: %v", err)
	}
}

// 更新统计信息
func (aa *AlertAggregator) updateStats(action string, count int) {
	aa.statsMutex.Lock()
	defer aa.statsMutex.Unlock()

	switch action {
	case "new_group":
		aa.stats.TotalGroups += count
		aa.stats.ActiveGroups += count
	case "add_alert":
		aa.stats.TotalAlerts += count
	case "flush_expired":
		aa.stats.FlushedGroups += count
		aa.stats.ActiveGroups -= count
	}
	
	// 计算平均组大小
	if aa.stats.TotalGroups > 0 {
		aa.stats.AverageGroupSize = float64(aa.stats.TotalAlerts) / float64(aa.stats.TotalGroups)
	}
}

// 获取统计信息
func (aa *AlertAggregator) GetStats() *AggregationStats {
	aa.statsMutex.RLock()
	defer aa.statsMutex.RUnlock()

	// 从数据库获取最新统计
	dbStats, err := GetAggregationStats()
	if err == nil {
		return dbStats
	}

	// 返回内存统计
	stats := *aa.stats
	aa.mutex.RLock()
	stats.ActiveGroups = len(aa.groups)
	aa.mutex.RUnlock()
	
	return &stats
}

// 获取活跃组数量
func (aa *AlertAggregator) GetActiveGroupCount() int {
	aa.mutex.RLock()
	defer aa.mutex.RUnlock()
	return len(aa.groups)
}

// 获取所有活跃组
func (aa *AlertAggregator) GetAllActiveGroups() map[string]*AggregationGroup {
	aa.mutex.RLock()
	defer aa.mutex.RUnlock()

	result := make(map[string]*AggregationGroup)
	for k, v := range aa.groups {
		result[k] = v
	}
	return result
}

// 手动刷新组
func (aa *AlertAggregator) FlushGroup(groupKey string) (*AggregatedAlert, error) {
	aa.mutex.Lock()
	defer aa.mutex.Unlock()

	group, exists := aa.groups[groupKey]
	if !exists {
		return nil, fmt.Errorf("聚合组不存在: %s", groupKey)
	}

	// 生成聚合消息
	aggregated := aa.GenerateAggregatedMessage(group)
	
	// 从活跃组中移除
	delete(aa.groups, groupKey)
	aa.updateStats("flush_manual", 1)
	
	// 持久化
	go aa.persistToDatabase(group)
	
	logs.Info("[Aggregator] 手动刷新聚合组: %s", groupKey)
	return aggregated, nil
}

// 清除所有组
func (aa *AlertAggregator) ClearAllGroups() {
	aa.mutex.Lock()
	defer aa.mutex.Unlock()
	
	count := len(aa.groups)
	aa.groups = make(map[string]*AggregationGroup)
	
	logs.Info("[Aggregator] 清除所有聚合组: %d 个", count)
}

// 检查是否启用
func (aa *AlertAggregator) IsEnabled() bool {
	return aa.config.Enabled
}

// 启用聚合功能
func (aa *AlertAggregator) Enable() {
	aa.config.Enabled = true
	if aa.flushTimer == nil {
		aa.startFlushTimer()
	}
	logs.Info("[Aggregator] 聚合功能已启用")
}

// 禁用聚合功能
func (aa *AlertAggregator) Disable() {
	aa.config.Enabled = false
	if aa.flushTimer != nil {
		aa.flushTimer.Stop()
		aa.flushTimer = nil
	}
	logs.Info("[Aggregator] 聚合功能已禁用")
}

// 重新加载配置
func (aa *AlertAggregator) ReloadConfig(config *AggregationConfig) {
	aa.mutex.Lock()
	defer aa.mutex.Unlock()
	
	oldEnabled := aa.config.Enabled
	aa.config = config
	
	// 处理启用状态变化
	if !oldEnabled && config.Enabled {
		aa.startFlushTimer()
	} else if oldEnabled && !config.Enabled {
		if aa.flushTimer != nil {
			aa.flushTimer.Stop()
			aa.flushTimer = nil
		}
	}
	
	logs.Info("[Aggregator] 配置已重新加载")
}

// 停止聚合管理器
func (aa *AlertAggregator) Stop() {
	if aa.flushTimer != nil {
		aa.flushTimer.Stop()
	}
	logs.Info("[Aggregator] 聚合管理器已停止")
}

// 导出状态
func (aa *AlertAggregator) ExportState() (string, error) {
	aa.mutex.RLock()
	defer aa.mutex.RUnlock()

	state := make(map[string]interface{})
	state["config"] = aa.config
	state["stats"] = aa.stats
	state["active_groups"] = len(aa.groups)
	
	// 导出组信息（不包含完整告警数据）
	groups := make(map[string]interface{})
	for key, group := range aa.groups {
		groups[key] = map[string]interface{}{
			"count":      group.Count,
			"first_seen": group.FirstSeen,
			"last_seen":  group.LastSeen,
			"severity":   group.Severity,
			"status":     group.Status,
		}
	}
	state["groups"] = groups

	data, err := json.Marshal(state)
	if err != nil {
		return "", err
	}
	return string(data), nil
}