package models

import (
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

type DbConfig struct {
	Base config.Configer
}

func NewDbConfig(base config.Configer) *DbConfig {
	return &DbConfig{Base: base}
}

// getString checks if a key is present in our DB key-value cache
func (d *DbConfig) getString(key string) (string, bool) {
	return GetCacheConfig(key)
}

func (d *DbConfig) Set(key, val string) error {
	SetCacheConfig(key, val)
	return nil
}

func (d *DbConfig) String(key string) string {
	if val, ok := d.getString(key); ok {
		return val
	}
	return d.Base.String(key)
}

func (d *DbConfig) Strings(key string) []string {
	if val, ok := d.getString(key); ok {
		return strings.Split(val, ";")
	}
	return d.Base.Strings(key)
}

func (d *DbConfig) Int(key string) (int, error) {
	if val, ok := d.getString(key); ok {
		return strconv.Atoi(val)
	}
	return d.Base.Int(key)
}

func (d *DbConfig) Int64(key string) (int64, error) {
	if val, ok := d.getString(key); ok {
		return strconv.ParseInt(val, 10, 64)
	}
	return d.Base.Int64(key)
}

func (d *DbConfig) Bool(key string) (bool, error) {
	if val, ok := d.getString(key); ok {
		return strconv.ParseBool(val)
	}
	return d.Base.Bool(key)
}

func (d *DbConfig) Float(key string) (float64, error) {
	if val, ok := d.getString(key); ok {
		return strconv.ParseFloat(val, 64)
	}
	return d.Base.Float(key)
}

func (d *DbConfig) DefaultString(key, defaultVal string) string {
	if val, ok := d.getString(key); ok {
		return val
	}
	return d.Base.DefaultString(key, defaultVal)
}

func (d *DbConfig) DefaultStrings(key string, defaultVal []string) []string {
	if val, ok := d.getString(key); ok {
		return strings.Split(val, ";")
	}
	return d.Base.DefaultStrings(key, defaultVal)
}

func (d *DbConfig) DefaultInt(key string, defaultVal int) int {
	if val, ok := d.getString(key); ok {
		if v, err := strconv.Atoi(val); err == nil {
			return v
		}
	}
	return d.Base.DefaultInt(key, defaultVal)
}

func (d *DbConfig) DefaultInt64(key string, defaultVal int64) int64 {
	if val, ok := d.getString(key); ok {
		if v, err := strconv.ParseInt(val, 10, 64); err == nil {
			return v
		}
	}
	return d.Base.DefaultInt64(key, defaultVal)
}

func (d *DbConfig) DefaultBool(key string, defaultVal bool) bool {
	if val, ok := d.getString(key); ok {
		if v, err := strconv.ParseBool(val); err == nil {
			return v
		}
	}
	return d.Base.DefaultBool(key, defaultVal)
}

func (d *DbConfig) DefaultFloat(key string, defaultVal float64) float64 {
	if val, ok := d.getString(key); ok {
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			return v
		}
	}
	return d.Base.DefaultFloat(key, defaultVal)
}

func (d *DbConfig) DIY(key string) (interface{}, error) {
	if val, ok := d.getString(key); ok {
		return val, nil
	}
	return d.Base.DIY(key)
}

func (d *DbConfig) GetSection(section string) (map[string]string, error) {
	return d.Base.GetSection(section)
}

func (d *DbConfig) SaveConfigFile(filename string) error {
	return d.Base.SaveConfigFile(filename)
}

// HookBeegoAppConfig uses reflection and unsafe to bypass unexported field limits in Beego
func HookBeegoAppConfig() {
	val := reflect.ValueOf(beego.AppConfig).Elem()
	field := val.FieldByName("innerConfig")

	// Get a pointer to the unexported field
	ptr := unsafe.Pointer(field.UnsafeAddr())

	// Read the original config.Configer directly via unsafe pointer to bypass reflect.Value.Interface restriction
	originalInnerConfig := *(*config.Configer)(ptr)

	// Create our DbConfig wrapping the original innerConfig instead of the outer AppConfig
	newDbConfig := NewDbConfig(originalInnerConfig)

	// Write our new DbConfig to the innerConfig field of beego.AppConfig
	*(*config.Configer)(ptr) = newDbConfig
}
