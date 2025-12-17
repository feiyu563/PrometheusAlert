package test

import (
	"PrometheusAlert/controllers"
	"PrometheusAlert/models"
	"testing"
)

// 测试告警处理器集成功能
func TestAlertProcessorIntegration(t *testing.T) {
	// 初始化配置管理器
	err := models.InitConfigManager()
	if err != nil {
		t.Fatalf("初始化配置管理器失败: %v", err)
	}
	
	// 初始化告警处理器
	controllers.InitAlertProcessor()
	processor := controllers.GetGlobalAlertProcessor()
	
	if processor == nil {
		t.Fatalf("获取告警处理器失败")
	}
	
	if !processor.IsEnabled() {
		t.Errorf("告警处理器应该是启用状态")
	}
	
	// 测试Prometheus告警处理
	promAlert := map[string]interface{}{
		"alertname": "HighCPU",
		"instance":  "server1:9100",
		"status":    "firing",
		"labels": map[string]interface{}{
			"job":      "node-exporter",
			"severity": "warning",
		},
		"annotations": map[string]interface{}{
			"summary":     "CPU usage is high",
			"description": "CPU usage is above 80%",
		},
	}
	
	// 第一次处理应该允许发送
	result1, err := processor.ProcessAlert(promAlert, "prometheus")
	if err != nil {
		t.Fatalf("第一次处理告警失败: %v", err)
	}
	
	if !result1.ShouldSend {
		t.Errorf("第一次告警应该被发送")
	}
	
	if result1.Action != "new" {
		t.Errorf("期望动作为 'new'，实际为 '%s'", result1.Action)
	}
	
	// 第二次处理相同告警应该被去重
	result2, err := processor.ProcessAlert(promAlert, "prometheus")
	if err != nil {
		t.Fatalf("第二次处理告警失败: %v", err)
	}
	
	if result2.ShouldSend {
		t.Errorf("第二次相同告警应该被去重")
	}
	
	if result2.Action != "duplicate" {
		t.Errorf("期望动作为 'duplicate'，实际为 '%s'", result2.Action)
	}
	
	if result2.Count != 2 {
		t.Errorf("期望计数为 2，实际为 %d", result2.Count)
	}
	
	// 检查缓存大小
	if processor.GetCacheSize() != 1 {
		t.Errorf("期望缓存大小为 1，实际为 %d", processor.GetCacheSize())
	}
	
	t.Logf("告警处理器集成测试通过")
}

// 测试阿里云告警处理
func TestAliyunAlertProcessing(t *testing.T) {
	// 初始化配置管理器
	err := models.InitConfigManager()
	if err != nil {
		t.Fatalf("初始化配置管理器失败: %v", err)
	}
	
	// 初始化告警处理器
	controllers.InitAlertProcessor()
	processor := controllers.GetGlobalAlertProcessor()
	
	// 清除缓存以确保测试独立性
	processor.ClearCache()
	
	// 测试阿里云告警
	aliyunAlert := map[string]interface{}{
		"alertName":    "基础监控-ECS-内存使用率",
		"instanceName": "instance-name-test",
		"metricName":   "Host.mem.usedutilization",
		"namespace":    "acs_ecs",
		"triggerLevel": "WARN",
		"alertState":   "ALERT",
		"curValue":     "97.39",
		"expression":   "$Average>=95",
		"userId":       "12345",
		"timestamp":    "1508136760",
	}
	
	// 处理阿里云告警
	result, err := processor.ProcessAlert(aliyunAlert, "aliyun")
	if err != nil {
		t.Fatalf("处理阿里云告警失败: %v", err)
	}
	
	if !result.ShouldSend {
		t.Errorf("阿里云告警应该被发送")
	}
	
	if result.Action != "new" {
		t.Errorf("期望动作为 'new'，实际为 '%s'", result.Action)
	}
	
	// 再次发送相同告警应该被去重
	result2, err := processor.ProcessAlert(aliyunAlert, "aliyun")
	if err != nil {
		t.Fatalf("第二次处理阿里云告警失败: %v", err)
	}
	
	if result2.ShouldSend {
		t.Errorf("重复的阿里云告警应该被去重")
	}
	
	t.Logf("阿里云告警处理测试通过")
}

// 测试不同来源告警的独立性
func TestMultiSourceAlerts(t *testing.T) {
	// 初始化配置管理器
	err := models.InitConfigManager()
	if err != nil {
		t.Fatalf("初始化配置管理器失败: %v", err)
	}
	
	// 初始化告警处理器
	controllers.InitAlertProcessor()
	processor := controllers.GetGlobalAlertProcessor()
	
	// 清除缓存
	processor.ClearCache()
	
	// Prometheus告警
	promAlert := map[string]interface{}{
		"alertname": "HighCPU",
		"instance":  "server1:9100",
		"status":    "firing",
		"labels": map[string]interface{}{
			"job":      "node-exporter",
			"severity": "warning",
		},
	}
	
	// 阿里云告警（相似但不同）
	aliyunAlert := map[string]interface{}{
		"alertName":    "HighCPU",
		"instanceName": "server1:9100",
		"metricName":   "Host.cpu.utilization",
		"namespace":    "acs_ecs",
		"triggerLevel": "WARN",
		"alertState":   "ALERT",
	}
	
	// 处理Prometheus告警
	result1, err := processor.ProcessAlert(promAlert, "prometheus")
	if err != nil {
		t.Fatalf("处理Prometheus告警失败: %v", err)
	}
	
	if !result1.ShouldSend {
		t.Errorf("Prometheus告警应该被发送")
	}
	
	// 处理阿里云告警（应该是独立的，不会被去重）
	result2, err := processor.ProcessAlert(aliyunAlert, "aliyun")
	if err != nil {
		t.Fatalf("处理阿里云告警失败: %v", err)
	}
	
	if !result2.ShouldSend {
		t.Errorf("不同来源的告警应该被独立处理")
	}
	
	// 缓存中应该有两个不同的告警
	if processor.GetCacheSize() != 2 {
		t.Errorf("期望缓存大小为 2，实际为 %d", processor.GetCacheSize())
	}
	
	t.Logf("多来源告警独立性测试通过")
}

// 测试告警处理器配置管理
func TestAlertProcessorConfig(t *testing.T) {
	// 初始化配置管理器
	err := models.InitConfigManager()
	if err != nil {
		t.Fatalf("初始化配置管理器失败: %v", err)
	}
	
	// 初始化告警处理器
	controllers.InitAlertProcessor()
	processor := controllers.GetGlobalAlertProcessor()
	
	// 测试启用/禁用功能
	if !processor.IsEnabled() {
		t.Errorf("告警处理器应该默认启用")
	}
	
	// 禁用处理器
	processor.Disable()
	if processor.IsEnabled() {
		t.Errorf("告警处理器应该被禁用")
	}
	
	// 测试禁用状态下的告警处理
	promAlert := map[string]interface{}{
		"alertname": "TestAlert",
		"instance":  "test:9100",
		"status":    "firing",
	}
	
	result, err := processor.ProcessAlert(promAlert, "prometheus")
	if err != nil {
		t.Fatalf("禁用状态下处理告警失败: %v", err)
	}
	
	if !result.ShouldSend {
		t.Errorf("禁用状态下告警应该直接通过")
	}
	
	if result.Action != "disabled" {
		t.Errorf("期望动作为 'disabled'，实际为 '%s'", result.Action)
	}
	
	// 重新启用
	processor.Enable()
	if !processor.IsEnabled() {
		t.Errorf("告警处理器应该被重新启用")
	}
	
	// 测试统计信息
	stats := processor.GetStats()
	if stats == nil {
		t.Errorf("获取统计信息失败")
	}
	
	t.Logf("告警处理器配置管理测试通过")
}