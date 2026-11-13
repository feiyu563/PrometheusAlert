package models

import (
	"github.com/astaxie/beego/orm"
)

type OnCall struct {
	Id    int64
	Date  string `orm:"size(30);unique"`
	Users string `orm:"type(text)"`
}

type OnCallUser struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func init() {
	orm.RegisterModel(new(OnCall))
}

func AddOnCall(oc OnCall) error {
	o := orm.NewOrm()
	_, err := o.Insert(&oc)
	return err
}

func UpdateOnCall(oc OnCall) error {
	o := orm.NewOrm()
	_, err := o.Update(&oc)
	return err
}

func GetOnCallById(id int64) (OnCall, error) {
	o := orm.NewOrm()
	oc := OnCall{Id: id}
	err := o.Read(&oc)
	return oc, err
}

func GetAllOnCall() ([]*OnCall, error) {
	o := orm.NewOrm()
	var ocs []*OnCall
	_, err := o.QueryTable("on_call").OrderBy("-date").All(&ocs)
	return ocs, err
}

func DelOnCall(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&OnCall{Id: id})
	return err
}

func GetOnCallByDateString(date string) (OnCall, error) {
	o := orm.NewOrm()
	oc := OnCall{Date: date}
	err := o.Read(&oc, "Date")
	return oc, err
}
