package controllers

import (
	"PrometheusAlert/models"
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// 去重管理控制器
type DeduplicationController struct {
	beego.Controller
}

// 去重统计页面
func (c *DeduplicationController) Stats() {
	processor := GetGlobalAlertProcessor()
	
	// 获取统计信息
	stats := processor.GetStats()
	cacheSize := processor.GetCacheSize()
	isEnabled := processor.IsEnabled()
	
	// 获取配置信息
	config := models.GetGlobalConfig()
	
	c.Data["Stats"] = stats
	c.Data["CacheSize"] = cacheSize
	c.Data["IsEnabled"] = isEnabled
	c.Data["Config"] = config.Deduplication
	c.Data["Title"] = "告警去重统计"
	
	c.TplName = "deduplication_stats.html"
}

// 获取统计信息API
func (c *DeduplicationController) GetStats() {
	processor := GetGlobalAlertProcessor()
	
	stats := processor.GetStats()
	cacheSize := processor.GetCacheSize()
	isEnabled := processor.IsEnabled()
	
	response := map[string]interface{}{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"stats":      stats,
			"cache_size": cacheSize,
			"enabled":    isEnabled,
		},
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 获取缓存的告警列表
func (c *DeduplicationController) GetCachedAlerts() {
	processor := GetGlobalAlertProcessor()
	cachedAlerts := processor.GetAllCachedAlerts()
	
	// 转换为前端友好的格式
	alertList := make([]map[string]interface{}, 0)
	for fingerprint, cached := range cachedAlerts {
		alert := map[string]interface{}{
			"fingerprint":    fingerprint,
			"alert_name":     cached.LastAlert.AlertName,
			"instance":       cached.LastAlert.Instance,
			"severity":       cached.LastAlert.Severity,
			"status":         cached.Status,
			"count":          cached.Count,
			"first_seen":     cached.FirstSeen.Format("2006-01-02 15:04:05"),
			"last_seen":      cached.LastSeen.Format("2006-01-02 15:04:05"),
			"suppress_until": "",
		}
		
		if !cached.SuppressUntil.IsZero() {
			alert["suppress_until"] = cached.SuppressUntil.Format("2006-01-02 15:04:05")
		}
		
		alertList = append(alertList, alert)
	}
	
	response := map[string]interface{}{
		"code":    200,
		"message": "success",
		"data":    alertList,
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 清除缓存
func (c *DeduplicationController) ClearCache() {
	processor := GetGlobalAlertProcessor()
	processor.ClearCache()
	
	logs.Info("[DeduplicationController] 用户清除了去重缓存")
	
	response := map[string]interface{}{
		"code":    200,
		"message": "缓存已清除",
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 启用/禁用去重功能
func (c *DeduplicationController) Toggle() {
	processor := GetGlobalAlertProcessor()
	
	enabledStr := c.GetString("enabled")
	enabled, err := strconv.ParseBool(enabledStr)
	if err != nil {
		response := map[string]interface{}{
			"code":    400,
			"message": "无效的参数",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	if enabled {
		processor.Enable()
	} else {
		processor.Disable()
	}
	
	logs.Info("[DeduplicationController] 用户%s了去重功能", map[bool]string{true: "启用", false: "禁用"}[enabled])
	
	response := map[string]interface{}{
		"code":    200,
		"message": map[bool]string{true: "去重功能已启用", false: "去重功能已禁用"}[enabled],
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 手动抑制告警
func (c *DeduplicationController) SuppressAlert() {
	fingerprint := c.GetString("fingerprint")
	duration := c.GetString("duration")
	
	if fingerprint == "" || duration == "" {
		response := map[string]interface{}{
			"code":    400,
			"message": "缺少必要参数",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	processor := GetGlobalAlertProcessor()
	err := processor.SuppressAlert(fingerprint, duration)
	if err != nil {
		logs.Error("[DeduplicationController] 抑制告警失败: %v", err)
		response := map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	logs.Info("[DeduplicationController] 用户手动抑制告警: %s, 持续时间: %s", fingerprint, duration)
	
	response := map[string]interface{}{
		"code":    200,
		"message": "告警已被抑制",
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 取消抑制告警
func (c *DeduplicationController) UnsuppressAlert() {
	fingerprint := c.GetString("fingerprint")
	
	if fingerprint == "" {
		response := map[string]interface{}{
			"code":    400,
			"message": "缺少fingerprint参数",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	processor := GetGlobalAlertProcessor()
	err := processor.UnsuppressAlert(fingerprint)
	if err != nil {
		logs.Error("[DeduplicationController] 取消抑制告警失败: %v", err)
		response := map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	logs.Info("[DeduplicationController] 用户取消抑制告警: %s", fingerprint)
	
	response := map[string]interface{}{
		"code":    200,
		"message": "已取消抑制",
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 重新加载配置
func (c *DeduplicationController) ReloadConfig() {
	// 重新加载全局配置
	err := models.GlobalConfigManager.LoadFromBeegoConfig()
	if err != nil {
		logs.Error("[DeduplicationController] 重新加载配置失败: %v", err)
		response := map[string]interface{}{
			"code":    500,
			"message": "重新加载配置失败: " + err.Error(),
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	// 重新加载告警处理器配置
	processor := GetGlobalAlertProcessor()
	err = processor.ReloadConfig()
	if err != nil {
		logs.Error("[DeduplicationController] 重新加载告警处理器配置失败: %v", err)
		response := map[string]interface{}{
			"code":    500,
			"message": "重新加载告警处理器配置失败: " + err.Error(),
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	logs.Info("[DeduplicationController] 用户重新加载了去重配置")
	
	response := map[string]interface{}{
		"code":    200,
		"message": "配置已重新加载",
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 获取配置信息
func (c *DeduplicationController) GetConfig() {
	config := models.GetGlobalConfig()
	
	response := map[string]interface{}{
		"code":    200,
		"message": "success",
		"data":    config,
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 更新配置
func (c *DeduplicationController) UpdateConfig() {
	var configData map[string]interface{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &configData)
	if err != nil {
		response := map[string]interface{}{
			"code":    400,
			"message": "无效的JSON格式",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	// 这里可以添加配置更新逻辑
	// 由于配置结构比较复杂，暂时只支持重新加载
	
	response := map[string]interface{}{
		"code":    501,
		"message": "配置更新功能待实现，请使用重新加载功能",
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 获取去重记录历史
func (c *DeduplicationController) GetHistory() {
	// 获取查询参数
	pageStr := c.GetString("page", "1")
	limitStr := c.GetString("limit", "20")
	alertName := c.GetString("alert_name")
	
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	
	var records []*models.AlertDeduplicationRecord
	var err error
	
	if alertName != "" {
		records, err = models.GetDeduplicationRecordsByAlertName(alertName)
	} else {
		records, err = models.GetAllDeduplicationRecords()
	}
	
	if err != nil {
		logs.Error("[DeduplicationController] 获取去重记录失败: %v", err)
		response := map[string]interface{}{
			"code":    500,
			"message": "获取记录失败",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	// 简单分页
	total := len(records)
	start := (page - 1) * limit
	end := start + limit
	
	if start >= total {
		records = []*models.AlertDeduplicationRecord{}
	} else {
		if end > total {
			end = total
		}
		records = records[start:end]
	}
	
	response := map[string]interface{}{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"records": records,
			"total":   total,
			"page":    page,
			"limit":   limit,
		},
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 导出缓存状态
func (c *DeduplicationController) ExportCache() {
	processor := GetGlobalAlertProcessor()
	
	if processor.deduplicator == nil {
		response := map[string]interface{}{
			"code":    500,
			"message": "去重管理器未初始化",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	cacheState, err := processor.deduplicator.ExportCacheState()
	if err != nil {
		logs.Error("[DeduplicationController] 导出缓存状态失败: %v", err)
		response := map[string]interface{}{
			"code":    500,
			"message": "导出失败: " + err.Error(),
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	// 设置下载头
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Ctx.Output.Header("Content-Disposition", "attachment; filename=dedup_cache_"+time.Now().Format("20060102_150405")+".json")
	
	c.Ctx.Output.Body([]byte(cacheState))
}

// 聚合管理控制器
type AggregationController struct {
	beego.Controller
}

// 获取聚合统计信息
func (c *AggregationController) GetStats() {
	processor := GetGlobalAlertProcessor()
	
	stats := processor.GetAggregationStats()
	activeGroups := processor.GetActiveGroupCount()
	isEnabled := processor.IsAggregationEnabled()
	
	response := map[string]interface{}{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"stats":         stats,
			"active_groups": activeGroups,
			"enabled":       isEnabled,
		},
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 获取活跃聚合组列表
func (c *AggregationController) GetActiveGroups() {
	processor := GetGlobalAlertProcessor()
	activeGroups := processor.GetAllActiveGroups()
	
	// 转换为前端友好的格式
	groupList := make([]map[string]interface{}, 0)
	for groupKey, group := range activeGroups {
		groupInfo := map[string]interface{}{
			"group_key":   groupKey,
			"count":       group.Count,
			"first_seen":  group.FirstSeen.Format("2006-01-02 15:04:05"),
			"last_seen":   group.LastSeen.Format("2006-01-02 15:04:05"),
			"severity":    group.Severity,
			"status":      group.Status,
			"labels":      group.Labels,
			"duration":    time.Since(group.FirstSeen).String(),
		}
		
		groupList = append(groupList, groupInfo)
	}
	
	response := map[string]interface{}{
		"code":    200,
		"message": "success",
		"data":    groupList,
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 手动刷新聚合组
func (c *AggregationController) FlushGroup() {
	groupKey := c.GetString("group_key")
	
	if groupKey == "" {
		response := map[string]interface{}{
			"code":    400,
			"message": "缺少group_key参数",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	processor := GetGlobalAlertProcessor()
	aggregatedAlert, err := processor.FlushAggregationGroup(groupKey)
	if err != nil {
		logs.Error("[AggregationController] 刷新聚合组失败: %v", err)
		response := map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	logs.Info("[AggregationController] 用户手动刷新聚合组: %s", groupKey)
	
	response := map[string]interface{}{
		"code":    200,
		"message": "聚合组已刷新",
		"data":    aggregatedAlert,
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 清除所有聚合组
func (c *AggregationController) ClearAllGroups() {
	processor := GetGlobalAlertProcessor()
	processor.ClearAllAggregationGroups()
	
	logs.Info("[AggregationController] 用户清除了所有聚合组")
	
	response := map[string]interface{}{
		"code":    200,
		"message": "所有聚合组已清除",
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 启用/禁用聚合功能
func (c *AggregationController) Toggle() {
	enabledStr := c.GetString("enabled")
	enabled, err := strconv.ParseBool(enabledStr)
	if err != nil {
		response := map[string]interface{}{
			"code":    400,
			"message": "无效的参数",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	processor := GetGlobalAlertProcessor()
	if enabled {
		processor.EnableAggregation()
	} else {
		processor.DisableAggregation()
	}
	
	logs.Info("[AggregationController] 用户%s了聚合功能", map[bool]string{true: "启用", false: "禁用"}[enabled])
	
	response := map[string]interface{}{
		"code":    200,
		"message": map[bool]string{true: "聚合功能已启用", false: "聚合功能已禁用"}[enabled],
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 获取聚合记录历史
func (c *AggregationController) GetHistory() {
	// 获取查询参数
	pageStr := c.GetString("page", "1")
	limitStr := c.GetString("limit", "20")
	
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	
	records, err := models.GetAllAggregationRecords()
	if err != nil {
		logs.Error("[AggregationController] 获取聚合记录失败: %v", err)
		response := map[string]interface{}{
			"code":    500,
			"message": "获取记录失败",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	// 简单分页
	total := len(records)
	start := (page - 1) * limit
	end := start + limit
	
	if start >= total {
		records = []*models.AlertAggregationRecord{}
	} else {
		if end > total {
			end = total
		}
		records = records[start:end]
	}
	
	response := map[string]interface{}{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"records": records,
			"total":   total,
			"page":    page,
			"limit":   limit,
		},
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 获取聚合记录详情
func (c *AggregationController) GetRecordDetail() {
	idStr := c.GetString("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response := map[string]interface{}{
			"code":    400,
			"message": "无效的ID参数",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	// 这里需要添加根据ID获取记录的方法
	response := map[string]interface{}{
		"code":    501,
		"message": "功能待实现",
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}

// 搜索聚合记录
func (c *AggregationController) SearchRecords() {
	keyword := c.GetString("keyword")
	if keyword == "" {
		response := map[string]interface{}{
			"code":    400,
			"message": "缺少搜索关键词",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	records, err := models.SearchAggregationRecords(keyword)
	if err != nil {
		logs.Error("[AggregationController] 搜索聚合记录失败: %v", err)
		response := map[string]interface{}{
			"code":    500,
			"message": "搜索失败",
		}
		c.Data["json"] = response
		c.ServeJSON()
		return
	}
	
	response := map[string]interface{}{
		"code":    200,
		"message": "success",
		"data":    records,
	}
	
	c.Data["json"] = response
	c.ServeJSON()
}