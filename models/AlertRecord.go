package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type AlertRecord struct {
	Id          int64
	Alertname   string
	AlertLevel  string
	Labels      string
	Instance    string
	StartsAt    string
	EndsAt      string
	Summary     string
	Description string
	AlertStatus string
	CreatedTime time.Time
	UpdatedBy   string
	UpdatedTime time.Time
}

func GetAllRecord() ([]*AlertRecord, error) {
	o := orm.NewOrm()
	Record_all := make([]*AlertRecord, 0)
	qs := o.QueryTable("AlertRecord")
	_, err := qs.OrderBy("-id").All(&Record_all)
	return Record_all, err
}

func GetRecordExist(alertname, alertLevel, lables, instance, startsAt, endsAt, summary, description, alertStatus string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("AlertRecord")
	flag := qs.Filter("Alertname", alertname).Filter("AlertLevel", alertLevel).Filter("Labels", lables).Filter("Instance", instance).Filter("Summary", summary).Filter("Description", description).Filter("StartsAt", startsAt).Filter("EndsAt", endsAt).Filter("AlertStatus", alertStatus).Exist()
	return flag
}

func RecordClean() {
	o := orm.NewOrm()
	o.Raw("delete from alert_record").Exec()
}

func RecordCleanByTime(RecordLiveDay int) {
	o := orm.NewOrm()
	o.Raw("delete from alert_record where created_time < ?", time.Now().AddDate(0, 0, RecordLiveDay*-1)).Exec()
}

func AddAlertRecord(alertname, alertLevel, labels, instance, startsAt, endsAt, summary, description, alertStatus string) error {
	o := orm.NewOrm()
	var err error

	alertRecord := &AlertRecord{
		//Id: id,
		Alertname:   alertname,
		AlertLevel:  alertLevel,
		Labels:      labels,
		Instance:    instance,
		StartsAt:    startsAt,
		EndsAt:      endsAt,
		Summary:     summary,
		Description: description,
		AlertStatus: alertStatus,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	// 插入数据
	_, err = o.Insert(alertRecord)
	return err
}
