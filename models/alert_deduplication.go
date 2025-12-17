package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/orm"
)

// 去重记录数据模型
type AlertDeduplicationRecord struct {
	Id            int64     `json:"id" orm:"auto"`
	Fingerprint   string    `json:"fingerprint" orm:"size(64);unique"`
	AlertName     string    `json:"alert_name" orm:"size(255)"`
	Instance      string    `json:"instance" orm:"size(255)"`
	Labels        string    `json:"labels" orm:"type(text)"`
	FirstSeen     time.Time `json:"first_seen"`
	LastSeen      time.Time `json:"last_seen"`
	Count         int       `json:"count" orm:"default(1)"`
	Status        string    `json:"status" orm:"size(20);default(active)"`
	SuppressUntil time.Time `json:"suppress_until" orm:"null"`
	CreatedAt     time.Time `json:"created_at" orm:"auto_now_add"`
	UpdatedAt     time.Time `json:"updated_at" orm:"auto_now"`
}

// 表名
func (adr *AlertDeduplicationRecord) TableName() string {
	return "alert_deduplication"
}

// 获取所有去重记录
func GetAllDeduplicationRecords() ([]*AlertDeduplicationRecord, error) {
	o := orm.NewOrm()
	records := make([]*AlertDeduplicationRecord, 0)
	qs := o.QueryTable("alert_deduplication")
	_, err := qs.OrderBy("-last_seen").All(&records)
	return records, err
}

// 根据指纹获取去重记录
func GetDeduplicationRecordByFingerprint(fingerprint string) (*AlertDeduplicationRecord, error) {
	o := orm.NewOrm()
	record := &AlertDeduplicationRecord{}
	qs := o.QueryTable("alert_deduplication")
	err := qs.Filter("fingerprint", fingerprint).One(record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// 检查去重记录是否存在
func DeduplicationRecordExists(fingerprint string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("alert_deduplication")
	return qs.Filter("fingerprint", fingerprint).Exist()
}

// 添加去重记录
func AddDeduplicationRecord(fingerprint, alertName, instance, labels string) error {
	o := orm.NewOrm()
	
	record := &AlertDeduplicationRecord{
		Fingerprint: fingerprint,
		AlertName:   alertName,
		Instance:    instance,
		Labels:      labels,
		FirstSeen:   time.Now(),
		LastSeen:    time.Now(),
		Count:       1,
		Status:      "active",
	}
	
	_, err := o.Insert(record)
	return err
}

// 更新去重记录
func UpdateDeduplicationRecord(fingerprint string, count int) error {
	o := orm.NewOrm()
	
	record := &AlertDeduplicationRecord{}
	qs := o.QueryTable("alert_deduplication")
	err := qs.Filter("fingerprint", fingerprint).One(record)
	if err != nil {
		return err
	}
	
	record.Count = count
	record.LastSeen = time.Now()
	record.UpdatedAt = time.Now()
	
	_, err = o.Update(record, "count", "last_seen", "updated_at")
	return err
}

// 设置抑制时间
func SetDeduplicationSuppressUntil(fingerprint string, suppressUntil time.Time) error {
	o := orm.NewOrm()
	
	record := &AlertDeduplicationRecord{}
	qs := o.QueryTable("alert_deduplication")
	err := qs.Filter("fingerprint", fingerprint).One(record)
	if err != nil {
		return err
	}
	
	record.SuppressUntil = suppressUntil
	record.UpdatedAt = time.Now()
	
	_, err = o.Update(record, "suppress_until", "updated_at")
	return err
}

// 删除过期的去重记录
func CleanExpiredDeduplicationRecords(expireDuration time.Duration) error {
	o := orm.NewOrm()
	expireTime := time.Now().Add(-expireDuration)
	
	_, err := o.Raw("DELETE FROM alert_deduplication WHERE last_seen < ?", expireTime).Exec()
	return err
}

// 获取去重统计信息
func GetDeduplicationStats() (*DeduplicationStats, error) {
	o := orm.NewOrm()
	
	stats := &DeduplicationStats{}
	
	// 总记录数
	totalCount, err := o.QueryTable("alert_deduplication").Count()
	if err != nil {
		return nil, err
	}
	stats.TotalRecords = int(totalCount)
	
	// 活跃记录数
	activeCount, err := o.QueryTable("alert_deduplication").Filter("status", "active").Count()
	if err != nil {
		return nil, err
	}
	stats.ActiveRecords = int(activeCount)
	
	// 今天的记录数
	today := time.Now().Truncate(24 * time.Hour)
	todayCount, err := o.QueryTable("alert_deduplication").Filter("created_at__gte", today).Count()
	if err != nil {
		return nil, err
	}
	stats.TodayRecords = int(todayCount)
	
	// 总去重次数
	var totalDuplicates int64
	err = o.Raw("SELECT SUM(count - 1) FROM alert_deduplication WHERE count > 1").QueryRow(&totalDuplicates)
	if err != nil {
		totalDuplicates = 0
	}
	stats.TotalDuplicates = int(totalDuplicates)
	
	return stats, nil
}

// 获取热门告警(去重次数最多)
func GetTopDuplicatedAlerts(limit int) ([]*AlertDeduplicationRecord, error) {
	o := orm.NewOrm()
	records := make([]*AlertDeduplicationRecord, 0)
	qs := o.QueryTable("alert_deduplication")
	_, err := qs.OrderBy("-count").Limit(limit).All(&records)
	return records, err
}

// 根据告警名称获取去重记录
func GetDeduplicationRecordsByAlertName(alertName string) ([]*AlertDeduplicationRecord, error) {
	o := orm.NewOrm()
	records := make([]*AlertDeduplicationRecord, 0)
	qs := o.QueryTable("alert_deduplication")
	_, err := qs.Filter("alert_name", alertName).OrderBy("-last_seen").All(&records)
	return records, err
}

// 根据时间范围获取去重记录
func GetDeduplicationRecordsByTimeRange(startTime, endTime time.Time) ([]*AlertDeduplicationRecord, error) {
	o := orm.NewOrm()
	records := make([]*AlertDeduplicationRecord, 0)
	qs := o.QueryTable("alert_deduplication")
	_, err := qs.Filter("created_at__gte", startTime).Filter("created_at__lte", endTime).OrderBy("-created_at").All(&records)
	return records, err
}

// 去重统计信息
type DeduplicationStats struct {
	TotalRecords    int `json:"total_records"`
	ActiveRecords   int `json:"active_records"`
	TodayRecords    int `json:"today_records"`
	TotalDuplicates int `json:"total_duplicates"`
}

// 转换为JSON
func (adr *AlertDeduplicationRecord) ToJSON() (string, error) {
	data, err := json.Marshal(adr)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// 从JSON创建
func NewDeduplicationRecordFromJSON(jsonStr string) (*AlertDeduplicationRecord, error) {
	record := &AlertDeduplicationRecord{}
	err := json.Unmarshal([]byte(jsonStr), record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// 获取标签映射
func (adr *AlertDeduplicationRecord) GetLabelsMap() (map[string]string, error) {
	if adr.Labels == "" {
		return make(map[string]string), nil
	}
	
	labels := make(map[string]string)
	err := json.Unmarshal([]byte(adr.Labels), &labels)
	if err != nil {
		return nil, err
	}
	return labels, nil
}

// 设置标签映射
func (adr *AlertDeduplicationRecord) SetLabelsMap(labels map[string]string) error {
	if len(labels) == 0 {
		adr.Labels = ""
		return nil
	}
	
	data, err := json.Marshal(labels)
	if err != nil {
		return err
	}
	adr.Labels = string(data)
	return nil
}

// 检查是否被抑制
func (adr *AlertDeduplicationRecord) IsSuppressed() bool {
	if adr.SuppressUntil.IsZero() {
		return false
	}
	return time.Now().Before(adr.SuppressUntil)
}

// 检查是否活跃
func (adr *AlertDeduplicationRecord) IsActive() bool {
	return adr.Status == "active"
}

// 获取持续时间
func (adr *AlertDeduplicationRecord) GetDuration() time.Duration {
	return adr.LastSeen.Sub(adr.FirstSeen)
}

// 获取去重率
func (adr *AlertDeduplicationRecord) GetDeduplicationRate() float64 {
	if adr.Count <= 1 {
		return 0.0
	}
	return float64(adr.Count-1) / float64(adr.Count) * 100
}