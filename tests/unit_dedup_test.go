package test

import (
	"PrometheusAlert/models"
	"testing"
	"time"
)

// 测试指纹生成（不依赖数据库）
func TestFingerprintOnly(t *testing.T) {
	fingerprintConfig := &models.FingerprintConfig{
		Algorithm:     "md5",
		IncludeFields: []string{"alert_name", "instance", "labels"},
		ExcludeLabels: []string{},
		IncludeLabels: []string{},
	}
	
	fingerprinter := models.NewAlertFingerprinter(fingerprintConfig)
	
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
		Instance:  "server1:9100",
		Labels: map[string]string{
			"job":      "node-exporter",
			"severity": "warning",
		},
		Status:   "firing",
		Severity: "warning",
		Summary:  "CPU usage is very high", // 不同的summary
		Source:   "prometheus",
	}
	
	fp1 := fingerprinter.GenerateFingerprint(alert1)
	fp2 := fingerprinter.GenerateFingerprint(alert2)
	
	// 相同的告警应该生成相同的指纹（summary不参与指纹计算）
	if fp1.Hash != fp2.Hash {
		t.Errorf("相同告警应该生成相同指纹，但得到不同指纹: %s vs %s", fp1.Hash, fp2.Hash)
	}
	
	// 验证指纹长度（MD5应该是32位）
	if len(fp1.Hash) != 32 {
		t.Errorf("MD5指纹长度应该为32，实际为 %d", len(fp1.Hash))
	}
	
	t.Logf("指纹生成测试通过，指纹: %s", fp1.Hash)
}

// 测试告警标准化（不依赖数据库）
func TestAlertNormalization(t *testing.T) {
	normalizer := models.NewAlertNormalizer()
	
	// 测试Prometheus告警转换
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
	
	standardAlert, err := normalizer.Normalize(promAlert, "prometheus")
	if err != nil {
		t.Fatalf("Prometheus告警标准化失败: %v", err)
	}
	
	if standardAlert.AlertName != "HighCPU" {
		t.Errorf("期望告警名为 'HighCPU'，实际为 '%s'", standardAlert.AlertName)
	}
	
	if standardAlert.Instance != "server1:9100" {
		t.Errorf("期望实例为 'server1:9100'，实际为 '%s'", standardAlert.Instance)
	}
	
	if standardAlert.Status != "firing" {
		t.Errorf("期望状态为 'firing'，实际为 '%s'", standardAlert.Status)
	}
	
	if standardAlert.Source != "prometheus" {
		t.Errorf("期望来源为 'prometheus'，实际为 '%s'", standardAlert.Source)
	}
	
	// 检查标签
	if len(standardAlert.Labels) != 2 {
		t.Errorf("期望标签数量为 2，实际为 %d", len(standardAlert.Labels))
	}
	
	if standardAlert.Labels["job"] != "node-exporter" {
		t.Errorf("期望job标签为 'node-exporter'，实际为 '%s'", standardAlert.Labels["job"])
	}
	
	t.Logf("Prometheus告警标准化测试通过")
}

// 测试阿里云告警转换（不依赖数据库）
func TestAliyunAlertNormalization(t *testing.T) {
	normalizer := models.NewAlertNormalizer()
	
	// 测试阿里云告警转换
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
	
	standardAlert, err := normalizer.Normalize(aliyunAlert, "aliyun")
	if err != nil {
		t.Fatalf("阿里云告警标准化失败: %v", err)
	}
	
	if standardAlert.AlertName != "基础监控-ECS-内存使用率" {
		t.Errorf("期望告警名为 '基础监控-ECS-内存使用率'，实际为 '%s'", standardAlert.AlertName)
	}
	
	if standardAlert.Instance != "instance-name-test" {
		t.Errorf("期望实例为 'instance-name-test'，实际为 '%s'", standardAlert.Instance)
	}
	
	if standardAlert.Status != "firing" {
		t.Errorf("期望状态为 'firing'，实际为 '%s'", standardAlert.Status)
	}
	
	if standardAlert.Severity != "warning" {
		t.Errorf("期望严重级别为 'warning'，实际为 '%s'", standardAlert.Severity)
	}
	
	if standardAlert.Source != "aliyun" {
		t.Errorf("期望来源为 'aliyun'，实际为 '%s'", standardAlert.Source)
	}
	
	// 检查标签
	if len(standardAlert.Labels) == 0 {
		t.Errorf("阿里云告警应该有标签")
	}
	
	if standardAlert.Labels["metricName"] != "Host.mem.usedutilization" {
		t.Errorf("期望metricName标签为 'Host.mem.usedutilization'，实际为 '%s'", standardAlert.Labels["metricName"])
	}
	
	t.Logf("阿里云告警标准化测试通过")
}

// 测试配置管理（不依赖数据库）
func TestConfigManagerOnly(t *testing.T) {
	configManager := models.NewConfigManager()
	
	// 测试默认配置
	config := configManager.GetConfig()
	if config == nil {
		t.Fatalf("获取配置失败")
	}
	
	if !config.Deduplication.Enabled {
		t.Errorf("默认配置应该启用去重功能")
	}
	
	if config.Deduplication.TimeWindow != 5*time.Minute {
		t.Errorf("默认时间窗口应该为5分钟，实际为 %v", config.Deduplication.TimeWindow)
	}
	
	if config.Deduplication.MaxCount != 5 {
		t.Errorf("默认最大计数应该为5，实际为 %d", config.Deduplication.MaxCount)
	}
	
	// 测试配置验证
	err := configManager.ValidateConfig()
	if err != nil {
		t.Errorf("默认配置验证失败: %v", err)
	}
	
	// 测试配置摘要
	summary := configManager.GetConfigSummary()
	if summary == nil {
		t.Errorf("获取配置摘要失败")
	}
	
	if !summary.DeduplicationEnabled {
		t.Errorf("配置摘要显示去重功能未启用")
	}
	
	// 测试JSON序列化
	jsonStr, err := configManager.ToJSON()
	if err != nil {
		t.Errorf("配置JSON序列化失败: %v", err)
	}
	
	if len(jsonStr) == 0 {
		t.Errorf("JSON序列化结果为空")
	}
	
	t.Logf("配置管理测试通过")
}

// 测试内存缓存（不依赖数据库）
func TestMemoryCacheOnly(t *testing.T) {
	cache := models.NewMemoryCache(100, 1*time.Hour)
	defer cache.Stop()
	
	// 测试设置和获取
	cache.Set("test_key", "test_value")
	
	value, exists := cache.Get("test_key")
	if !exists {
		t.Errorf("缓存项应该存在")
	}
	
	if value != "test_value" {
		t.Errorf("期望值为 'test_value'，实际为 '%v'", value)
	}
	
	// 测试缓存大小
	if cache.Size() != 1 {
		t.Errorf("期望缓存大小为 1，实际为 %d", cache.Size())
	}
	
	// 测试存在性检查
	if !cache.Exists("test_key") {
		t.Errorf("缓存项应该存在")
	}
	
	// 测试Peek（不更新访问时间）
	peekValue, exists := cache.Peek("test_key")
	if !exists {
		t.Errorf("Peek应该找到缓存项")
	}
	
	if peekValue != "test_value" {
		t.Errorf("Peek期望值为 'test_value'，实际为 '%v'", peekValue)
	}
	
	// 测试删除
	deleted := cache.Delete("test_key")
	if !deleted {
		t.Errorf("删除操作应该成功")
	}
	
	if cache.Size() != 0 {
		t.Errorf("删除后缓存大小应该为 0，实际为 %d", cache.Size())
	}
	
	// 测试批量操作
	items := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	
	cache.SetBatch(items)
	
	if cache.Size() != 3 {
		t.Errorf("批量设置后缓存大小应该为 3，实际为 %d", cache.Size())
	}
	
	batchResult := cache.GetBatch([]string{"key1", "key2", "key4"})
	if len(batchResult) != 2 {
		t.Errorf("批量获取应该返回 2 个项，实际为 %d", len(batchResult))
	}
	
	// 测试统计信息
	stats := cache.GetStats()
	if stats == nil {
		t.Errorf("获取统计信息失败")
	}
	
	if stats.Size != 3 {
		t.Errorf("统计信息显示缓存大小应该为 3，实际为 %d", stats.Size)
	}
	
	t.Logf("内存缓存测试通过")
}

// 测试策略管理（不依赖数据库）
func TestPolicyManager(t *testing.T) {
	policyManager := models.NewPolicyManager()
	
	// 测试获取所有策略
	policies := policyManager.GetAllPolicies()
	if len(policies) == 0 {
		t.Errorf("应该有内置策略")
	}
	
	// 测试获取特定策略
	strictPolicy, exists := policyManager.GetPolicy("strict")
	if !exists {
		t.Errorf("应该存在严格策略")
	}
	
	if strictPolicy.GetName() != "strict" {
		t.Errorf("策略名称应该为 'strict'，实际为 '%s'", strictPolicy.GetName())
	}
	
	// 测试宽松策略
	loosePolicy, exists := policyManager.GetPolicy("loose")
	if !exists {
		t.Errorf("应该存在宽松策略")
	}
	
	if loosePolicy.GetName() != "loose" {
		t.Errorf("策略名称应该为 'loose'，实际为 '%s'", loosePolicy.GetName())
	}
	
	// 测试基于严重级别的策略
	severityPolicy, exists := policyManager.GetPolicy("severity_based")
	if !exists {
		t.Errorf("应该存在基于严重级别的策略")
	}
	
	if severityPolicy.GetName() != "severity_based" {
		t.Errorf("策略名称应该为 'severity_based'，实际为 '%s'", severityPolicy.GetName())
	}
	
	t.Logf("策略管理测试通过，共有 %d 个策略", len(policies))
}