package models

import (
	"sync"

	"github.com/astaxie/beego/orm"
)

type SysConfig struct {
	Id    int
	Key   string `orm:"unique;index"`
	Value string `orm:"type(text)"`
}

var (
	ConfigCache   = make(map[string]string)
	ConfigCacheMu sync.RWMutex
)

func AddSysConfig(key, value string) error {
	o := orm.NewOrm()
	config := &SysConfig{Key: key, Value: value}
	_, err := o.Insert(config)
	return err
}

func UpdateSysConfig(key, value string) error {
	o := orm.NewOrm()
	config := &SysConfig{Key: key}
	err := o.Read(config, "Key")
	if err == nil {
		config.Value = value
		_, err = o.Update(config, "Value")
		return err
	} else if err == orm.ErrNoRows {
		config.Value = value
		_, err = o.Insert(config)
		return err
	}
	return err
}

func GetSysConfig(key string) (string, error) {
	o := orm.NewOrm()
	config := &SysConfig{Key: key}
	err := o.Read(config, "Key")
	if err != nil {
		return "", err
	}
	return config.Value, nil
}

func GetAllSysConfig() (map[string]string, error) {
	o := orm.NewOrm()
	var list []*SysConfig
	_, err := o.QueryTable("SysConfig").All(&list)
	if err != nil {
		return nil, err
	}
	res := make(map[string]string)
	for _, item := range list {
		res[item.Key] = item.Value
	}
	return res, nil
}

func InitConfigCache() {
	list, err := GetAllSysConfig()
	if err == nil {
		ConfigCacheMu.Lock()
		ConfigCache = list
		ConfigCacheMu.Unlock()
	}
}

func GetCacheConfig(key string) (string, bool) {
	ConfigCacheMu.RLock()
	val, ok := ConfigCache[key]
	ConfigCacheMu.RUnlock()
	return val, ok
}

func SetCacheConfig(key, val string) {
	ConfigCacheMu.Lock()
	ConfigCache[key] = val
	ConfigCacheMu.Unlock()
	_ = UpdateSysConfig(key, val)
}
