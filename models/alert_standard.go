package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// 标准化告警结构
type StandardAlert struct {
	AlertName   string            `json:"alert_name"`
	Instance    string            `json:"instance"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	Status      string            `json:"status"`      // firing, resolved
	Severity    string            `json:"severity"`    // critical, warning, info
	StartsAt    time.Time         `json:"starts_at"`
	EndsAt      time.Time         `json:"ends_at"`
	Summary     string            `json:"summary"`
	Description string            `json:"description"`
	Source      string            `json:"source"`      // prometheus, zabbix, aliyun
	RawData     interface{}       `json:"raw_data"`    // 原始数据
}

// 聚合告警结构
type AggregatedAlert struct {
	GroupKey     string           `json:"group_key"`
	Count        int              `json:"count"`
	FirstSeen    time.Time        `json:"first_seen"`
	LastSeen     time.Time        `json:"last_seen"`
	Alerts       []*StandardAlert `json:"alerts"`
	Summary      string           `json:"summary"`
	Description  string           `json:"description"`
	IsAggregated bool             `json:"is_aggregated"`
	Severity     string           `json:"severity"`
	Status       string           `json:"status"`
}

// 告警转换器接口
type AlertConverter interface {
	Convert(rawAlert interface{}) (*StandardAlert, error)
	GetSource() string
	Validate(rawAlert interface{}) error
}

// Prometheus告警转换器
type PrometheusAlertConverter struct{}

func (pac *PrometheusAlertConverter) GetSource() string {
	return "prometheus"
}

func (pac *PrometheusAlertConverter) Validate(rawAlert interface{}) error {
	alertMap, ok := rawAlert.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid prometheus alert format: not a map")
	}
	
	if _, exists := alertMap["alertname"]; !exists {
		return fmt.Errorf("missing required field: alertname")
	}
	
	return nil
}

func (pac *PrometheusAlertConverter) Convert(rawAlert interface{}) (*StandardAlert, error) {
	if err := pac.Validate(rawAlert); err != nil {
		return nil, err
	}
	
	alertMap := rawAlert.(map[string]interface{})
	
	return &StandardAlert{
		AlertName:   pac.getString(alertMap, "alertname"),
		Instance:    pac.getString(alertMap, "instance"),
		Labels:      pac.getLabels(alertMap, "labels"),
		Annotations: pac.getLabels(alertMap, "annotations"),
		Status:      pac.getString(alertMap, "status"),
		Severity:    pac.getSeverity(alertMap),
		StartsAt:    pac.getTime(alertMap, "startsAt"),
		EndsAt:      pac.getTime(alertMap, "endsAt"),
		Summary:     pac.getAnnotation(alertMap, "summary"),
		Description: pac.getAnnotation(alertMap, "description"),
		Source:      pac.GetSource(),
		RawData:     rawAlert,
	}, nil
}

func (pac *PrometheusAlertConverter) getString(alertMap map[string]interface{}, key string) string {
	if value, exists := alertMap[key]; exists {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}

func (pac *PrometheusAlertConverter) getLabels(alertMap map[string]interface{}, key string) map[string]string {
	labels := make(map[string]string)
	if labelsInterface, exists := alertMap[key]; exists {
		if labelsMap, ok := labelsInterface.(map[string]interface{}); ok {
			for k, v := range labelsMap {
				if str, ok := v.(string); ok {
					labels[k] = str
				}
			}
		}
	}
	return labels
}

func (pac *PrometheusAlertConverter) getSeverity(alertMap map[string]interface{}) string {
	// 从labels中获取severity
	if labelsInterface, exists := alertMap["labels"]; exists {
		if labelsMap, ok := labelsInterface.(map[string]interface{}); ok {
			if severity, exists := labelsMap["severity"]; exists {
				if str, ok := severity.(string); ok {
					return str
				}
			}
		}
	}
	return "info" // 默认级别
}

func (pac *PrometheusAlertConverter) getAnnotation(alertMap map[string]interface{}, key string) string {
	if annotationsInterface, exists := alertMap["annotations"]; exists {
		if annotationsMap, ok := annotationsInterface.(map[string]interface{}); ok {
			if value, exists := annotationsMap[key]; exists {
				if str, ok := value.(string); ok {
					return str
				}
			}
		}
	}
	return ""
}

func (pac *PrometheusAlertConverter) getTime(alertMap map[string]interface{}, key string) time.Time {
	if timeInterface, exists := alertMap[key]; exists {
		if timeStr, ok := timeInterface.(string); ok {
			if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
				return t
			}
		}
	}
	return time.Now()
}

// 阿里云告警转换器
type AliyunAlertConverter struct{}

func (aac *AliyunAlertConverter) GetSource() string {
	return "aliyun"
}

func (aac *AliyunAlertConverter) Validate(rawAlert interface{}) error {
	alertMap, ok := rawAlert.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid aliyun alert format: not a map")
	}
	
	if _, exists := alertMap["alertName"]; !exists {
		return fmt.Errorf("missing required field: alertName")
	}
	
	return nil
}

func (aac *AliyunAlertConverter) Convert(rawAlert interface{}) (*StandardAlert, error) {
	if err := aac.Validate(rawAlert); err != nil {
		return nil, err
	}
	
	alertMap := rawAlert.(map[string]interface{})
	
	labels := make(map[string]string)
	labels["metricName"] = aac.getString(alertMap, "metricName")
	labels["namespace"] = aac.getString(alertMap, "namespace")
	labels["userId"] = aac.getString(alertMap, "userId")
	
	// 解析dimensions
	if dimensions := aac.getString(alertMap, "dimensions"); dimensions != "" {
		// 简单解析dimensions字符串，实际可能需要更复杂的解析
		labels["dimensions"] = dimensions
	}
	
	return &StandardAlert{
		AlertName:   aac.getString(alertMap, "alertName"),
		Instance:    aac.getString(alertMap, "instanceName"),
		Labels:      labels,
		Annotations: make(map[string]string),
		Status:      aac.convertStatus(aac.getString(alertMap, "alertState")),
		Severity:    aac.convertSeverity(aac.getString(alertMap, "triggerLevel")),
		StartsAt:    aac.getTimestamp(alertMap, "timestamp"),
		EndsAt:      time.Time{}, // 阿里云告警没有结束时间
		Summary:     fmt.Sprintf("阿里云监控告警: %s", aac.getString(alertMap, "alertName")),
		Description: fmt.Sprintf("指标: %s, 当前值: %s, 表达式: %s", 
			aac.getString(alertMap, "metricName"),
			aac.getString(alertMap, "curValue"),
			aac.getString(alertMap, "expression")),
		Source:  aac.GetSource(),
		RawData: rawAlert,
	}, nil
}

func (aac *AliyunAlertConverter) getString(alertMap map[string]interface{}, key string) string {
	if value, exists := alertMap[key]; exists {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}

func (aac *AliyunAlertConverter) convertStatus(alertState string) string {
	switch alertState {
	case "ALERT":
		return "firing"
	case "OK":
		return "resolved"
	default:
		return "firing"
	}
}

func (aac *AliyunAlertConverter) convertSeverity(triggerLevel string) string {
	switch triggerLevel {
	case "CRITICAL":
		return "critical"
	case "WARN":
		return "warning"
	case "INFO":
		return "info"
	default:
		return "warning"
	}
}

func (aac *AliyunAlertConverter) getTimestamp(alertMap map[string]interface{}, key string) time.Time {
	if timestampInterface, exists := alertMap[key]; exists {
		if timestampStr, ok := timestampInterface.(string); ok {
			// 阿里云时间戳是Unix时间戳
			if timestamp, err := time.Parse("1136239445", timestampStr); err == nil {
				return timestamp
			}
		}
	}
	return time.Now()
}

// 告警标准化器
type AlertNormalizer struct {
	converters map[string]AlertConverter
}

func NewAlertNormalizer() *AlertNormalizer {
	normalizer := &AlertNormalizer{
		converters: make(map[string]AlertConverter),
	}
	
	// 注册转换器
	normalizer.RegisterConverter(&PrometheusAlertConverter{})
	normalizer.RegisterConverter(&AliyunAlertConverter{})
	
	return normalizer
}

func (an *AlertNormalizer) RegisterConverter(converter AlertConverter) {
	an.converters[converter.GetSource()] = converter
}

func (an *AlertNormalizer) Normalize(rawAlert interface{}, source string) (*StandardAlert, error) {
	converter, exists := an.converters[source]
	if !exists {
		return nil, fmt.Errorf("unsupported alert source: %s", source)
	}
	
	return converter.Convert(rawAlert)
}

func (an *AlertNormalizer) GetSupportedSources() []string {
	sources := make([]string, 0, len(an.converters))
	for source := range an.converters {
		sources = append(sources, source)
	}
	return sources
}

// 工具函数
func (sa *StandardAlert) ToJSON() (string, error) {
	data, err := json.Marshal(sa)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (sa *StandardAlert) GetLabelsString() string {
	if len(sa.Labels) == 0 {
		return ""
	}
	
	data, err := json.Marshal(sa.Labels)
	if err != nil {
		return ""
	}
	return string(data)
}

func (sa *StandardAlert) IsFiring() bool {
	return sa.Status == "firing"
}

func (sa *StandardAlert) IsResolved() bool {
	return sa.Status == "resolved"
}

func (aa *AggregatedAlert) ToJSON() (string, error) {
	data, err := json.Marshal(aa)
	if err != nil {
		return "", err
	}
	return string(data), nil
}