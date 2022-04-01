package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

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
