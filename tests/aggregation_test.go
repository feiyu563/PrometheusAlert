package test

import (
	"PrometheusAlert/models"
	"fmt"
	"testing"
	"time"
)

// 测试聚合管理器基本功能
func TestAggregatorBasic(t *testing.T) {
	config := &models.AggregationConfig{
		Enabled:       true,
		TimeWindow:    1 * time.Minute,
		MaxAlerts:     3,
		GroupByLabels: []string{"alertname", "severity"},
		Strategy:      "summary",
		FlushInterval: 10 * time.Second,
	}
	
	aggregator := models.NewAlertAggregator(config)
	defer aggregator.Stop()
	
	// 创建测试告警
	alert1 := &models.StandardAlert{
		AlertName: "HighCPU",
		Instance:  "server1:9100",
		Labels: map[string]string{
			"job":      "node-exporter",
			"severity": "warning",
		},
		Status:   "firing",
		Severity: "warning",
		Summary:  "CPU usage is high",
		Source:   "prometheus",
	}
	
	alert2 := &models.StandardAlert{
		AlertName: "HighCPU",
		Instance:  "server2:9100",
		Labels: map[string]string{
			"job":      "node-exporter",
			"severity": "warning",
		},
		Status:   "firing",
		Severity: "warning",
		Summary:  "CPU usage is high",
		Source:   "prometheus",
	}
	
	// 第一个告警应该创建新组
	result1, err := aggregator.AddAlert(alert1)
	if err != nil {
		t.Fatalf("添加第一个告警失败: %v", err)
	}
	
	if result1.ShouldFlush {
		t.Errorf("第一个告警不应该触发刷新")
	}
	
	if result1.Action != "aggregated" {
		t.Errorf("期望动作为 'aggregated'，实际为 '%s'", result1.Action)
	}
	
	if result1.Group.Count != 1 {
		t.Errorf("期望组计数为 1，实际为 %d", result1.Group.Count)
	}
	
	// 第二个告警应该加入同一组
	result2, err := aggregator.AddAlert(alert2)
	if err != nil {
		t.Fatalf("添加第二个告警失败: %v", err)
	}
	
	if result2.ShouldFlush {
		t.Errorf("第二个告警不应该触发刷新")
	}
	
	if result2.Group.Count != 2 {
		t.Errorf("期望组计数为 2，实际为 %d", result2.Group.Count)
	}
	
	// 检查活跃组数量
	if aggregator.GetActiveGroupCount() != 1 {
		t.Errorf("期望活跃组数量为 1，实际为 %d", aggregator.GetActiveGroupCount())
	}
	
	t.Logf("聚合管理器基本功能测试通过")
}

// 测试最大告警数触发刷新
func TestAggregatorMaxAlerts(t *testing.T) {
	config := &models.AggregationConfig{
		Enabled:       true,
		TimeWindow:    5 * time.Minute,
		MaxAlerts:     2, // 设置较小的最大值
		GroupByLabels: []string{"alertname"},
		Strategy:      "summary",
		FlushInterval: 10 * time.Second,
	}
	
	aggregator := models.NewAlertAggregator(config)
	defer aggregator.Stop()
	
	alert := &models.StandardAlert{
		AlertName: "HighCPU",
		Instance:  "server1:9100",
		Labels: map[string]string{
			"job":      "node-exporter",
			"severity": "warning",
		},
		Status:   "firing",
		Severity: "warning",
		Summary:  "CPU usage is high",
		Source:   "prometheus",
	}
	
	// 添加第一个告警
	result1, err := aggregator.AddAlert(alert)
	if err != nil {
		t.Fatalf("添加第一个告警失败: %v", err)
	}
	
	if result1.ShouldFlush {
		t.Errorf("第一个告警不应该触发刷新")
	}
	
	// 添加第二个告警，应该触发刷新
	result2, err := aggregator.AddAlert(alert)
	if err != nil {
		t.Fatalf("添加第二个告警失败: %v", err)
	}
	
	if !result2.ShouldFlush {
		t.Errorf("达到最大告警数应该触发刷新")
	}
	
	if result2.Group.Count != 2 {
		t.Errorf("期望组计数为 2，实际为 %d", result2.Group.Count)
	}
	
	t.Logf("最大告警数触发刷新测试通过")
}

// 测试分组键生成
func TestAggregatorGroupKey(t *testing.T) {
	config := &models.AggregationConfig{
		Enabled:       true,
		TimeWindow:    1 * time.Minute,
		MaxAlerts:     10,
		GroupByLabels: []string{"alertname", "severity", "labels.job"},
		Strategy:      "summary",
		FlushInterval: 10 * time.Second,
	}
	
	aggregator := models.NewAlertAggregator(config)
	defer aggregator.Stop()
	
	alert1 := &models.StandardAlert{
		AlertName: "HighCPU",
		Instance:  "server1:9100",
		Labels: map[string]string{
			"job":      "node-exporter",
			"severity": "warning",
		},
		Status:   "firing",
		Severity: "warning",
		Summary:  "CPU usage is high",
		Source:   "prometheus",
	}
	
	alert2 := &models.StandardAlert{
		AlertName: "HighCPU",
		Instance:  "server2:9100",
		Labels: map[string]string{
			"job":      "prometheus", // 不同的job
			"severity": "warning",
		},
		Status:   "firing",
		Severity: "warning",
		Summary:  "CPU usage is high",
		Source:   "prometheus",
	}
	
	// 添加第一个告警
	result1, err := aggregator.AddAlert(alert1)
	if err != nil {
		t.Fatalf("添加第一个告警失败: %v", err)
	}
	
	// 添加第二个告警（不同的job，应该创建新组）
	result2, err := aggregator.AddAlert(alert2)
	if err != nil {
		t.Fatalf("添加第二个告警失败: %v", err)
	}
	
	// 应该有两个不同的组
	if aggregator.GetActiveGroupCount() != 2 {
		t.Errorf("期望活跃组数量为 2，实际为 %d", aggregator.GetActiveGroupCount())
	}
	
	// 组键应该不同
	if result1.Group.GroupKey == result2.Group.GroupKey {
		t.Errorf("不同的告警应该生成不同的组键")
	}
	
	t.Logf("分组键生成测试通过")
}

// 测试聚合消息生成
func TestAggregatedMessageGeneration(t *testing.T) {
	config := &models.AggregationConfig{
		Enabled:       true,
		TimeWindow:    1 * time.Minute,
		MaxAlerts:     10,
		GroupByLabels: []string{"alertname"},
		Strategy:      "summary",
		FlushInterval: 10 * time.Second,
	}
	
	aggregator := models.NewAlertAggregator(config)
	defer aggregator.Stop()
	
	// 创建聚合组
	group := &models.AggregationGroup{
		GroupKey:  "alertname=HighCPU",
		Alerts:    make([]*models.StandardAlert, 0),
		FirstSeen: time.Now().Add(-5 * time.Minute),
		LastSeen:  time.Now(),
		Count:     0,
		Status:    "active",
		Severity:  "warning",
		Labels: map[string]string{
			"alertname": "HighCPU",
		},
	}
	
	// 添加一些告警
	for i := 0; i < 3; i++ {
		alert := &models.StandardAlert{
			AlertName: "HighCPU",
			Instance:  fmt.Sprintf("server%d:9100", i+1),
			Labels: map[string]string{
				"job":      "node-exporter",
				"severity": "warning",
			},
			Status:   "firing",
			Severity: "warning",
			Summary:  "CPU usage is high",
			Source:   "prometheus",
		}
		group.Alerts = append(group.Alerts, alert)
	}
	group.Count = len(group.Alerts)
	
	// 生成聚合消息
	aggregated := aggregator.GenerateAggregatedMessage(group)
	
	if aggregated == nil {
		t.Fatalf("生成聚合消息失败")
	}
	
	if !aggregated.IsAggregated {
		t.Errorf("聚合消息应该标记为已聚合")
	}
	
	if aggregated.Count != 3 {
		t.Errorf("期望聚合消息计数为 3，实际为 %d", aggregated.Count)
	}
	
	if aggregated.GroupKey != group.GroupKey {
		t.Errorf("聚合消息组键不匹配")
	}
	
	if aggregated.Summary == "" {
		t.Errorf("聚合消息摘要不应该为空")
	}
	
	if aggregated.Description == "" {
		t.Errorf("聚合消息描述不应该为空")
	}
	
	t.Logf("聚合消息生成测试通过")
	t.Logf("摘要: %s", aggregated.Summary)
	t.Logf("描述: %s", aggregated.Description)
}

// 测试不同聚合策略
func TestAggregationStrategies(t *testing.T) {
	strategies := []string{"count", "list", "summary"}
	
	for _, strategy := range strategies {
		t.Run(fmt.Sprintf("Strategy_%s", strategy), func(t *testing.T) {
			config := &models.AggregationConfig{
				Enabled:       true,
				TimeWindow:    1 * time.Minute,
				MaxAlerts:     10,
				GroupByLabels: []string{"alertname"},
				Strategy:      strategy,
				FlushInterval: 10 * time.Second,
			}
			
			aggregator := models.NewAlertAggregator(config)
			defer aggregator.Stop()
			
			// 创建测试组
			group := &models.AggregationGroup{
				GroupKey:  "alertname=HighCPU",
				Alerts:    make([]*models.StandardAlert, 0),
				FirstSeen: time.Now().Add(-2 * time.Minute),
				LastSeen:  time.Now(),
				Count:     2,
				Status:    "active",
				Severity:  "warning",
			}
			
			// 添加测试告警
			for i := 0; i < 2; i++ {
				alert := &models.StandardAlert{
					AlertName: "HighCPU",
					Instance:  fmt.Sprintf("server%d:9100", i+1),
					Status:    "firing",
					Severity:  "warning",
					Summary:   "CPU usage is high",
					Source:    "prometheus",
				}
				group.Alerts = append(group.Alerts, alert)
			}
			
			// 生成聚合消息
			aggregated := aggregator.GenerateAggregatedMessage(group)
			
			if aggregated == nil {
				t.Fatalf("策略 %s 生成聚合消息失败", strategy)
			}
			
			if aggregated.Summary == "" {
				t.Errorf("策略 %s 的摘要不应该为空", strategy)
			}
			
			if aggregated.Description == "" {
				t.Errorf("策略 %s 的描述不应该为空", strategy)
			}
			
			t.Logf("策略 %s 测试通过", strategy)
			t.Logf("摘要: %s", aggregated.Summary)
		})
	}
}

// 测试聚合功能启用/禁用
func TestAggregatorEnableDisable(t *testing.T) {
	config := &models.AggregationConfig{
		Enabled:       false, // 初始禁用
		TimeWindow:    1 * time.Minute,
		MaxAlerts:     10,
		GroupByLabels: []string{"alertname"},
		Strategy:      "summary",
		FlushInterval: 10 * time.Second,
	}
	
	aggregator := models.NewAlertAggregator(config)
	defer aggregator.Stop()
	
	if aggregator.IsEnabled() {
		t.Errorf("聚合功能应该初始禁用")
	}
	
	alert := &models.StandardAlert{
		AlertName: "HighCPU",
		Instance:  "server1:9100",
		Status:    "firing",
		Severity:  "warning",
		Summary:   "CPU usage is high",
		Source:    "prometheus",
	}
	
	// 禁用状态下添加告警
	result, err := aggregator.AddAlert(alert)
	if err != nil {
		t.Fatalf("禁用状态下添加告警失败: %v", err)
	}
	
	if !result.ShouldFlush {
		t.Errorf("禁用状态下告警应该直接刷新")
	}
	
	if result.Action != "disabled" {
		t.Errorf("期望动作为 'disabled'，实际为 '%s'", result.Action)
	}
	
	// 启用聚合功能
	aggregator.Enable()
	if !aggregator.IsEnabled() {
		t.Errorf("聚合功能应该被启用")
	}
	
	// 启用状态下添加告警
	result2, err := aggregator.AddAlert(alert)
	if err != nil {
		t.Fatalf("启用状态下添加告警失败: %v", err)
	}
	
	if result2.ShouldFlush {
		t.Errorf("启用状态下第一个告警不应该触发刷新")
	}
	
	if result2.Action != "aggregated" {
		t.Errorf("期望动作为 'aggregated'，实际为 '%s'", result2.Action)
	}
	
	t.Logf("聚合功能启用/禁用测试通过")
}

// 测试统计信息
func TestAggregatorStats(t *testing.T) {
	config := &models.AggregationConfig{
		Enabled:       true,
		TimeWindow:    1 * time.Minute,
		MaxAlerts:     10,
		GroupByLabels: []string{"alertname"},
		Strategy:      "summary",
		FlushInterval: 10 * time.Second,
	}
	
	aggregator := models.NewAlertAggregator(config)
	defer aggregator.Stop()
	
	alert := &models.StandardAlert{
		AlertName: "HighCPU",
		Instance:  "server1:9100",
		Status:    "firing",
		Severity:  "warning",
		Summary:   "CPU usage is high",
		Source:    "prometheus",
	}
	
	// 添加一些告警
	for i := 0; i < 3; i++ {
		aggregator.AddAlert(alert)
	}
	
	stats := aggregator.GetStats()
	if stats == nil {
		t.Fatalf("获取统计信息失败")
	}
	
	if stats.ActiveGroups != 1 {
		t.Errorf("期望活跃组数为 1，实际为 %d", stats.ActiveGroups)
	}
	
	t.Logf("聚合统计信息测试通过")
}