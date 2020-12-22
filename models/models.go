package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

// 分类
type PrometheusAlertDB struct {
	Id      int
	Tpltype string
	Tpluse  string
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
