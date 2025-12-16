package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/orm"
)

// 聚合记录数据模型
type AlertAggregationRecord struct {
	Id          int64     `json:"id" orm:"auto"`
	GroupKey    string    `json:"group_key" orm:"size(128)"`
	GroupLabels string    `json:"group_labels" orm:"type(text)"`
	AlertCount  int       `json:"alert_count" orm:"default(0)"`
	FirstSeen   time.Time `json:"first_seen"`
	LastSeen    time.Time `json:"last_seen"`
	Status      string    `json:"status" orm:"size(20);default(active)"`
	Summary     string    `json:"summary" orm:"type(text)"`
	Description string    `json:"description" orm:"type(text)"`
	AlertsData  string    `json:"alerts_data" orm:"type(text)"`
	CreatedAt   time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt   time.Time `json:"updated_at" orm:"auto_now"`
}

// 表名
func (aar *AlertAggregationRecord) TableName() string {
	return "alert_aggregation"
}

// 获取所有聚合记录
func GetAllAggregationRecords() ([]*AlertAggregationRecord, error) {
	o := orm.NewOrm()
	records := make([]*AlertAggregationRecord, 0)
	qs := o.QueryTable("alert_aggregation")
	_, err := qs.OrderBy("-last_seen").All(&records)
	return records, err
}

// 根据分组键获取聚合记录
func GetAggregationRecordByGroupKey(groupKey string) (*AlertAggregationRecord, error) {
	o := orm.NewOrm()
	record := &AlertAggregationRecord{}
	qs := o.QueryTable("alert_aggregation")
	err := qs.Filter("group_key", groupKey).One(record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// 检查聚合记录是否存在
func AggregationRecordExists(groupKey string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("alert_aggregation")
	return qs.Filter("group_key", groupKey).Exist()
}

// 添加聚合记录
func AddAggregationRecord(groupKey, groupLabels string, alertCount int, firstSeen, lastSeen time.Time, status, summary, description, alertsData string) error {
	o := orm.NewOrm()
	
	record := &AlertAggregationRecord{
		GroupKey:    groupKey,
		GroupLabels: groupLabels,
		AlertCount:  alertCount,
		FirstSeen:   firstSeen,
		LastSeen:    lastSeen,
		Status:      status,
		Summary:     summary,
		Description: description,
		AlertsData:  alertsData,
	}
	
	_, err := o.Insert(record)
	return err
}

// 更新聚合记录
func UpdateAggregationRecord(groupKey string, alertCount int, lastSeen time.Time, summary, description, alertsData string) error {
	o := orm.NewOrm()
	
	record := &AlertAggregationRecord{}
	qs := o.QueryTable("alert_aggregation")
	err := qs.Filter("group_key", groupKey).One(record)
	if err != nil {
		return err
	}
	
	record.AlertCount = alertCount
	record.LastSeen = lastSeen
	record.Summary = summary
	record.Description = description
	record.AlertsData = alertsData
	record.UpdatedAt = time.Now()
	
	_, err = o.Update(record, "alert_count", "last_seen", "summary", "description", "alerts_data", "updated_at")
	return err
}

// 删除过期的聚合记录
func CleanExpiredAggregationRecords(expireDuration time.Duration) error {
	o := orm.NewOrm()
	expireTime := time.Now().Add(-expireDuration)
	
	_, err := o.Raw("DELETE FROM alert_aggregation WHERE last_seen < ?", expireTime).Exec()
	return err
}

// 获取聚合统计信息
func GetAggregationStats() (*AggregationStats, error) {
	o := orm.NewOrm()
	
	stats := &AggregationStats{}
	
	// 总记录数
	totalCount, err := o.QueryTable("alert_aggregation").Count()
	if err != nil {
		return nil, err
	}
	stats.TotalGroups = int(totalCount)
	
	// 活跃记录数
	activeCount, err := o.QueryTable("alert_aggregation").Filter("status", "active").Count()
	if err != nil {
		return nil, err
	}
	stats.ActiveGroups = int(activeCount)
	
	// 已刷新记录数
	flushedCount, err := o.QueryTable("alert_aggregation").Filter("status", "flushed").Count()
	if err != nil {
		return nil, err
	}
	stats.FlushedGroups = int(flushedCount)
	
	// 总告警数
	var totalAlerts int64
	err = o.Raw("SELECT SUM(alert_count) FROM alert_aggregation").QueryRow(&totalAlerts)
	if err != nil {
		totalAlerts = 0
	}
	stats.TotalAlerts = int(totalAlerts)
	
	// 平均组大小
	if stats.TotalGroups > 0 {
		stats.AverageGroupSize = float64(stats.TotalAlerts) / float64(stats.TotalGroups)
	}
	
	return stats, nil
}

// 获取热门聚合组(告警数最多)
func GetTopAggregationGroups(limit int) ([]*AlertAggregationRecord, error) {
	o := orm.NewOrm()
	records := make([]*AlertAggregationRecord, 0)
	qs := o.QueryTable("alert_aggregation")
	_, err := qs.OrderBy("-alert_count").Limit(limit).All(&records)
	return records, err
}

// 根据时间范围获取聚合记录
func GetAggregationRecordsByTimeRange(startTime, endTime time.Time) ([]*AlertAggregationRecord, error) {
	o := orm.NewOrm()
	records := make([]*AlertAggregationRecord, 0)
	qs := o.QueryTable("alert_aggregation")
	_, err := qs.Filter("created_at__gte", startTime).Filter("created_at__lte", endTime).OrderBy("-created_at").All(&records)
	return records, err
}

// 根据状态获取聚合记录
func GetAggregationRecordsByStatus(status string) ([]*AlertAggregationRecord, error) {
	o := orm.NewOrm()
	records := make([]*AlertAggregationRecord, 0)
	qs := o.QueryTable("alert_aggregation")
	_, err := qs.Filter("status", status).OrderBy("-last_seen").All(&records)
	return records, err
}

// 转换为JSON
func (aar *AlertAggregationRecord) ToJSON() (string, error) {
	data, err := json.Marshal(aar)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// 从JSON创建
func NewAggregationRecordFromJSON(jsonStr string) (*AlertAggregationRecord, error) {
	record := &AlertAggregationRecord{}
	err := json.Unmarshal([]byte(jsonStr), record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// 获取分组标签映射
func (aar *AlertAggregationRecord) GetGroupLabelsMap() (map[string]string, error) {
	if aar.GroupLabels == "" {
		return make(map[string]string), nil
	}
	
	labels := make(map[string]string)
	err := json.Unmarshal([]byte(aar.GroupLabels), &labels)
	if err != nil {
		return nil, err
	}
	return labels, nil
}

// 设置分组标签映射
func (aar *AlertAggregationRecord) SetGroupLabelsMap(labels map[string]string) error {
	if len(labels) == 0 {
		aar.GroupLabels = ""
		return nil
	}
	
	data, err := json.Marshal(labels)
	if err != nil {
		return err
	}
	aar.GroupLabels = string(data)
	return nil
}

// 获取告警数据
func (aar *AlertAggregationRecord) GetAlertsData() ([]*StandardAlert, error) {
	if aar.AlertsData == "" {
		return make([]*StandardAlert, 0), nil
	}
	
	var alerts []*StandardAlert
	err := json.Unmarshal([]byte(aar.AlertsData), &alerts)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}

// 设置告警数据
func (aar *AlertAggregationRecord) SetAlertsData(alerts []*StandardAlert) error {
	if len(alerts) == 0 {
		aar.AlertsData = ""
		return nil
	}
	
	data, err := json.Marshal(alerts)
	if err != nil {
		return err
	}
	aar.AlertsData = string(data)
	return nil
}

// 检查是否活跃
func (aar *AlertAggregationRecord) IsActive() bool {
	return aar.Status == "active"
}

// 检查是否已刷新
func (aar *AlertAggregationRecord) IsFlushed() bool {
	return aar.Status == "flushed"
}

// 获取持续时间
func (aar *AlertAggregationRecord) GetDuration() time.Duration {
	return aar.LastSeen.Sub(aar.FirstSeen)
}

// 获取平均告警间隔
func (aar *AlertAggregationRecord) GetAverageInterval() time.Duration {
	if aar.AlertCount <= 1 {
		return 0
	}
	
	duration := aar.GetDuration()
	return time.Duration(int64(duration) / int64(aar.AlertCount-1))
}

// 聚合记录摘要
type AggregationRecordSummary struct {
	Id          int64     `json:"id"`
	GroupKey    string    `json:"group_key"`
	AlertCount  int       `json:"alert_count"`
	FirstSeen   time.Time `json:"first_seen"`
	LastSeen    time.Time `json:"last_seen"`
	Duration    string    `json:"duration"`
	Status      string    `json:"status"`
	Summary     string    `json:"summary"`
}

// 获取聚合记录摘要
func (aar *AlertAggregationRecord) GetSummary() *AggregationRecordSummary {
	duration := aar.GetDuration()
	
	return &AggregationRecordSummary{
		Id:         aar.Id,
		GroupKey:   aar.GroupKey,
		AlertCount: aar.AlertCount,
		FirstSeen:  aar.FirstSeen,
		LastSeen:   aar.LastSeen,
		Duration:   duration.String(),
		Status:     aar.Status,
		Summary:    aar.Summary,
	}
}

// 批量获取聚合记录摘要
func GetAggregationRecordSummaries(limit, offset int) ([]*AggregationRecordSummary, error) {
	records, err := GetAllAggregationRecords()
	if err != nil {
		return nil, err
	}
	
	summaries := make([]*AggregationRecordSummary, 0)
	
	start := offset
	end := offset + limit
	
	if start >= len(records) {
		return summaries, nil
	}
	
	if end > len(records) {
		end = len(records)
	}
	
	for i := start; i < end; i++ {
		summaries = append(summaries, records[i].GetSummary())
	}
	
	return summaries, nil
}

// 搜索聚合记录
func SearchAggregationRecords(keyword string) ([]*AlertAggregationRecord, error) {
	o := orm.NewOrm()
	records := make([]*AlertAggregationRecord, 0)
	
	// 在group_key和summary中搜索
	_, err := o.Raw("SELECT * FROM alert_aggregation WHERE group_key LIKE ? OR summary LIKE ? ORDER BY last_seen DESC", 
		"%"+keyword+"%", "%"+keyword+"%").QueryRows(&records)
	
	return records, err
}