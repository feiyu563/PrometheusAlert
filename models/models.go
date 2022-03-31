package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

// 分类
type PrometheusAlertDB struct {
	Id      int
	Tpltype string //发送类型如钉钉、企业微信、飞书等
	Tpluse  string //接受目标如Prometheus、WebHook、graylog
	Tplname string `orm:"index"`
	Tpl     string `orm:"type(text)"`
	Created time.Time
}

func GetAllTpl() ([]*PrometheusAlertDB, error) {
	o := orm.NewOrm()
	Tpl_all := make([]*PrometheusAlertDB, 0)
	qs := o.QueryTable("PrometheusAlertDB")
	_, err := qs.All(&Tpl_all)
	return Tpl_all, err
}

func GetTpl(id int) (*PrometheusAlertDB, error) {
	o := orm.NewOrm()
	tpl_one := new(PrometheusAlertDB)
	qs := o.QueryTable("PrometheusAlertDB")
	err := qs.Filter("id", id).One(tpl_one)
	if err != nil {
		return nil, err
	}
	return tpl_one, err
}

func GetTplOne(name string) (*PrometheusAlertDB, error) {
	o := orm.NewOrm()
	tpl_one := new(PrometheusAlertDB)
	qs := o.QueryTable("PrometheusAlertDB")
	err := qs.Filter("Tplname", name).One(tpl_one)
	if err != nil {
		return tpl_one, err
	}
	return tpl_one, err
}

func GetPromtheusTpl() ([]*PrometheusAlertDB, error) {
	o := orm.NewOrm()
	tpl := make([]*PrometheusAlertDB, 0)
	qs := o.QueryTable("PrometheusAlertDB")
	_, err := qs.Filter("tpluse", "Prometheus").All(&tpl)
	if err != nil {
		return nil, err
	}
	return tpl, err
}

func DelTpl(id int) error {
	o := orm.NewOrm()
	tpl_one := &PrometheusAlertDB{Id: id}
	_, err := o.Delete(tpl_one)
	return err
}

func AddTpl(id int, tplname, t_type, t_use, tpl string) error {
	o := orm.NewOrm()
	qs := o.QueryTable("PrometheusAlertDB")
	bExist := qs.Filter("Tplname", tplname).Exist()
	var err error
	if bExist {
		err = errors.New("模版名称已经存在！")
		return err
	}
	Template_table := &PrometheusAlertDB{
		Id:      id,
		Tplname: tplname,
		Tpltype: t_type,
		Tpluse:  t_use,
		Tpl:     tpl,
		Created: time.Now(),
	}
	// 插入数据
	_, err = o.Insert(Template_table)
	return err
}

func UpdateTpl(id int, tplname, t_type, t_use, tpl string) error {
	o := orm.NewOrm()
	tpl_update := &PrometheusAlertDB{Id: id}
	err := o.Read(tpl_update)
	if err == nil {
		tpl_update.Id = id
		tpl_update.Tplname = tplname
		tpl_update.Tpltype = t_type
		tpl_update.Tpluse = t_use
		tpl_update.Tpl = tpl
		tpl_update.Created = time.Now()
		_, err := o.Update(tpl_update)
		return err
	}
	return err
}

type AlertRecord struct {
	Id          int64
	Alertname   string
	AlertLevel  string
	Job         string
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

func (alertRecord *AlertRecord) TableName() string {
	return "alert_record"
}

func GetAllRecord() ([]*AlertRecord, error) {
	o := orm.NewOrm()
	Record_all := make([]*AlertRecord, 0)
	qs := o.QueryTable("AlertRecord")
	_, err := qs.All(&Record_all)
	return Record_all, err
}

func GetRecordExist(alertname, alertLevel, instance, job, startsAt, endsAt, summary, description, alertStatus string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("AlertRecord")
	flag := qs.Filter("Alertname", alertname).Filter("AlertLevel", alertLevel).Filter("Job", job).Filter("Instance", instance).Filter("Summary", summary).Filter("Description", description).Filter("StartsAt", startsAt).Filter("EndsAt", endsAt).Filter("AlertStatus", alertStatus).Exist()
	return flag
}

func RecordClean() {
	//o := orm.NewOrm()
	//var r orm.RawSeter
	//r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")

	o := orm.NewOrm()
	o.Raw("delete from alert_record").Exec()
}

func RecordCleanByTime(RecordLiveDay int) {
	o := orm.NewOrm()
	o.Raw("delete from alert_record where created_time < ?", time.Now().AddDate(0, 0, RecordLiveDay*-1)).Exec()
}

func AddAlertRecord(alertname, alertLevel, instance, job, startsAt, endsAt, summary, description, alertStatus string) error {
	o := orm.NewOrm()
	var err error

	alertRecord := &AlertRecord{
		//Id: id,
		Alertname:   alertname,
		AlertLevel:  alertLevel,
		Instance:    instance,
		Job:         job,
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

type AlertRouter struct {
	Id         int `orm:"index"`
	Name       string
	Tpl        *PrometheusAlertDB `orm:"rel(fk)"`
	Rules      string
	UrlOrPhone string
	AtSomeOne  string
	Created    time.Time
}

func AddAlertRouter(id int, tplid int, name, rules, url_or_phone, at_some_one string) error {
	tpl, _ := GetTpl(tplid)
	o := orm.NewOrm()
	AlertRouter_table := &AlertRouter{
		Id:         id,
		Name:       name,
		Tpl:        tpl,
		Rules:      rules,
		UrlOrPhone: url_or_phone,
		AtSomeOne:  at_some_one,
		Created:    time.Now(),
	}
	// 插入数据
	_, err := o.Insert(AlertRouter_table)
	return err
}

func UpdateAlertRouter(id int, tplid int, name, rules, url_or_phone, at_some_one string) error {
	tpl, _ := GetTpl(tplid)
	o := orm.NewOrm()
	router_update := &AlertRouter{Id: id}
	err := o.Read(router_update)
	if err == nil {
		router_update.Id = id
		router_update.Name = name
		router_update.Tpl = tpl
		router_update.Rules = rules
		router_update.UrlOrPhone = url_or_phone
		router_update.AtSomeOne = at_some_one
		router_update.Created = time.Now()
		_, err := o.Update(router_update)
		return err
	}
	return err
}

func DelAlertRouter(id int) error {
	o := orm.NewOrm()
	tpl_one := &AlertRouter{Id: id}
	_, err := o.Delete(tpl_one)
	return err
}

func GetAllAlertRouter() ([]*AlertRouter, error) {
	o := orm.NewOrm()
	Tpl_all := make([]*AlertRouter, 0)
	qs := o.QueryTable("AlertRouter")
	_, err := qs.RelatedSel().All(&Tpl_all)
	return Tpl_all, err
}

func GetAlertRouter(id int) (*AlertRouter, error) {
	o := orm.NewOrm()
	tpl_one := new(AlertRouter)
	qs := o.QueryTable("AlertRouter")
	err := qs.Filter("id", id).RelatedSel().One(tpl_one)
	if err != nil {
		return nil, err
	}
	return tpl_one, err
}
