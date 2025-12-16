package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

// 去重策略接口
type DeduplicationPolicy interface {
	ShouldSend(cached *CachedAlert, alert *StandardAlert, config *DeduplicationConfig) *DeduplicationResult
	GetName() string
	GetDescription() string
}

// 去重规则
type DeduplicationRule struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Enabled     bool              `json:"enabled"`
	Priority    int               `json:"priority"`
	Conditions  []RuleCondition   `json:"conditions"`
	Actions     []RuleAction      `json:"actions"`
	Labels      map[string]string `json:"labels"` // 匹配的标签
}

// 规则条件
type RuleCondition struct {
	Field    string `json:"field"`    // alert_name, severity, instance, labels.xxx
	Operator string `json:"operator"` // eq, ne, contains, regex, gt, lt
	Value    string `json:"value"`
	Regex    *regexp.Regexp `json:"-"` // 编译后的正则表达式
}

// 规则动作
type RuleAction struct {
	Type       string        `json:"type"`        // suppress, allow, modify_count, set_ttl
	Duration   time.Duration `json:"duration"`    // 抑制时长
	MaxCount   int           `json:"max_count"`   // 最大计数
	Interval   time.Duration `json:"interval"`    // 发送间隔
	Message    string        `json:"message"`     // 自定义消息
}

// 严格策略
type StrictPolicy struct{}

func (sp *StrictPolicy) GetName() string {
	return "strict"
}

func (sp *StrictPolicy) GetDescription() string {
	return "严格策略：只有第一次和状态变化时发送"
}

func (sp *StrictPolicy) ShouldSend(cached *CachedAlert, alert *StandardAlert, config *DeduplicationConfig) *DeduplicationResult {
	// 状态变化时发送
	if cached.LastAlert != nil && cached.LastAlert.Status != alert.Status {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "status_changed",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     fmt.Sprintf("告警状态从 %s 变为 %s", cached.LastAlert.Status, alert.Status),
		}
	}

	// 超过最大计数时抑制
	if cached.Count > config.MaxCount {
		cached.SuppressUntil = time.Now().Add(config.TimeWindow)
		return &DeduplicationResult{
			ShouldSend: false,
			Action:     "max_count_exceeded",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     fmt.Sprintf("超过最大重复次数 %d，抑制发送", config.MaxCount),
		}
	}

	// 重复告警不发送
	return &DeduplicationResult{
		ShouldSend: false,
		Action:     "duplicate",
		Count:      cached.Count,
		Cached:     cached,
		Reason:     fmt.Sprintf("重复告警，第 %d 次出现", cached.Count),
	}
}

// 宽松策略
type LoosePolicy struct{}

func (lp *LoosePolicy) GetName() string {
	return "loose"
}

func (lp *LoosePolicy) GetDescription() string {
	return "宽松策略：允许一定频率的重复发送"
}

func (lp *LoosePolicy) ShouldSend(cached *CachedAlert, alert *StandardAlert, config *DeduplicationConfig) *DeduplicationResult {
	// 状态变化时发送
	if cached.LastAlert != nil && cached.LastAlert.Status != alert.Status {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "status_changed",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     "告警状态变化",
		}
	}

	// 每隔一定次数发送一次
	sendInterval := config.MaxCount / 2
	if sendInterval < 1 {
		sendInterval = 1
	}

	if cached.Count%sendInterval == 0 {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "interval_send",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     fmt.Sprintf("间隔发送，第 %d 次", cached.Count),
		}
	}

	return &DeduplicationResult{
		ShouldSend: false,
		Action:     "duplicate",
		Count:      cached.Count,
		Cached:     cached,
		Reason:     "重复告警，等待间隔发送",
	}
}

// 基于严重级别的策略
type SeverityBasedPolicy struct{}

func (sbp *SeverityBasedPolicy) GetName() string {
	return "severity_based"
}

func (sbp *SeverityBasedPolicy) GetDescription() string {
	return "基于严重级别的策略：不同级别采用不同的去重策略"
}

func (sbp *SeverityBasedPolicy) ShouldSend(cached *CachedAlert, alert *StandardAlert, config *DeduplicationConfig) *DeduplicationResult {
	// 状态变化时发送
	if cached.LastAlert != nil && cached.LastAlert.Status != alert.Status {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "status_changed",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     "告警状态变化",
		}
	}

	// 根据严重级别调整策略
	var maxCount int
	var sendInterval int

	switch alert.Severity {
	case "critical":
		maxCount = config.MaxCount * 3     // 严重告警允许更多重复
		sendInterval = 1                   // 每次都发送
	case "warning":
		maxCount = config.MaxCount
		sendInterval = 2                   // 每2次发送一次
	case "info":
		maxCount = config.MaxCount / 2     // 信息告警减少重复
		sendInterval = 5                   // 每5次发送一次
	default:
		maxCount = config.MaxCount
		sendInterval = 3
	}

	if maxCount < 1 {
		maxCount = 1
	}
	if sendInterval < 1 {
		sendInterval = 1
	}

	// 检查是否超过最大次数
	if cached.Count > maxCount {
		cached.SuppressUntil = time.Now().Add(config.TimeWindow)
		return &DeduplicationResult{
			ShouldSend: false,
			Action:     "severity_max_exceeded",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     fmt.Sprintf("超过严重级别 %s 的最大次数 %d", alert.Severity, maxCount),
		}
	}

	// 按间隔发送
	if cached.Count%sendInterval == 0 {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "severity_interval",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     fmt.Sprintf("基于严重级别 %s 的间隔发送", alert.Severity),
		}
	}

	return &DeduplicationResult{
		ShouldSend: false,
		Action:     "severity_duplicate",
		Count:      cached.Count,
		Cached:     cached,
		Reason:     fmt.Sprintf("严重级别 %s 的重复告警", alert.Severity),
	}
}

// 自定义规则策略
type CustomRulePolicy struct {
	rules []*DeduplicationRule
}

func NewCustomRulePolicy(rules []*DeduplicationRule) *CustomRulePolicy {
	// 编译正则表达式
	for _, rule := range rules {
		for i := range rule.Conditions {
			if rule.Conditions[i].Operator == "regex" {
				if regex, err := regexp.Compile(rule.Conditions[i].Value); err == nil {
					rule.Conditions[i].Regex = regex
				} else {
					logs.Error("[CustomRulePolicy] 编译正则表达式失败: %s, %v", rule.Conditions[i].Value, err)
				}
			}
		}
	}

	return &CustomRulePolicy{rules: rules}
}

func (crp *CustomRulePolicy) GetName() string {
	return "custom_rule"
}

func (crp *CustomRulePolicy) GetDescription() string {
	return "自定义规则策略：基于用户定义的规则进行去重"
}

func (crp *CustomRulePolicy) ShouldSend(cached *CachedAlert, alert *StandardAlert, config *DeduplicationConfig) *DeduplicationResult {
	// 状态变化时总是发送
	if cached.LastAlert != nil && cached.LastAlert.Status != alert.Status {
		return &DeduplicationResult{
			ShouldSend: true,
			Action:     "status_changed",
			Count:      cached.Count,
			Cached:     cached,
			Reason:     "告警状态变化",
		}
	}

	// 按优先级处理规则
	for _, rule := range crp.rules {
		if !rule.Enabled {
			continue
		}

		if crp.matchRule(rule, alert) {
			return crp.executeRule(rule, cached, alert, config)
		}
	}

	// 没有匹配的规则，使用默认策略
	return crp.defaultBehavior(cached, alert, config)
}

// 匹配规则
func (crp *CustomRulePolicy) matchRule(rule *DeduplicationRule, alert *StandardAlert) bool {
	// 检查标签匹配
	if len(rule.Labels) > 0 {
		for key, value := range rule.Labels {
			if alertValue, exists := alert.Labels[key]; !exists || alertValue != value {
				return false
			}
		}
	}

	// 检查条件匹配
	for _, condition := range rule.Conditions {
		if !crp.matchCondition(&condition, alert) {
			return false
		}
	}

	return true
}

// 匹配条件
func (crp *CustomRulePolicy) matchCondition(condition *RuleCondition, alert *StandardAlert) bool {
	fieldValue := crp.getFieldValue(condition.Field, alert)

	switch condition.Operator {
	case "eq":
		return fieldValue == condition.Value
	case "ne":
		return fieldValue != condition.Value
	case "contains":
		return strings.Contains(fieldValue, condition.Value)
	case "regex":
		if condition.Regex != nil {
			return condition.Regex.MatchString(fieldValue)
		}
		return false
	case "gt":
		return fieldValue > condition.Value
	case "lt":
		return fieldValue < condition.Value
	default:
		return false
	}
}

// 获取字段值
func (crp *CustomRulePolicy) getFieldValue(field string, alert *StandardAlert) string {
	switch field {
	case "alert_name":
		return alert.AlertName
	case "severity":
		return alert.Severity
	case "instance":
		return alert.Instance
	case "status":
		return alert.Status
	case "source":
		return alert.Source
	case "summary":
		return alert.Summary
	case "description":
		return alert.Description
	default:
		// 检查是否是标签字段
		if strings.HasPrefix(field, "labels.") {
			labelKey := strings.TrimPrefix(field, "labels.")
			if value, exists := alert.Labels[labelKey]; exists {
				return value
			}
		}
		return ""
	}
}

// 执行规则
func (crp *CustomRulePolicy) executeRule(rule *DeduplicationRule, cached *CachedAlert, alert *StandardAlert, config *DeduplicationConfig) *DeduplicationResult {
	for _, action := range rule.Actions {
		switch action.Type {
		case "suppress":
			cached.SuppressUntil = time.Now().Add(action.Duration)
			return &DeduplicationResult{
				ShouldSend: false,
				Action:     "rule_suppressed",
				Count:      cached.Count,
				Cached:     cached,
				Reason:     fmt.Sprintf("规则 %s 抑制告警 %v", rule.Name, action.Duration),
			}

		case "allow":
			return &DeduplicationResult{
				ShouldSend: true,
				Action:     "rule_allowed",
				Count:      cached.Count,
				Cached:     cached,
				Reason:     fmt.Sprintf("规则 %s 允许发送", rule.Name),
			}

		case "modify_count":
			if cached.Count > action.MaxCount {
				return &DeduplicationResult{
					ShouldSend: false,
					Action:     "rule_count_exceeded",
					Count:      cached.Count,
					Cached:     cached,
					Reason:     fmt.Sprintf("规则 %s 超过最大次数 %d", rule.Name, action.MaxCount),
				}
			}

		case "interval":
			if action.Interval > 0 && time.Since(cached.LastSeen) < action.Interval {
				return &DeduplicationResult{
					ShouldSend: false,
					Action:     "rule_interval",
					Count:      cached.Count,
					Cached:     cached,
					Reason:     fmt.Sprintf("规则 %s 间隔时间未到", rule.Name),
				}
			}
		}
	}

	// 默认允许发送
	return &DeduplicationResult{
		ShouldSend: true,
		Action:     "rule_default",
		Count:      cached.Count,
		Cached:     cached,
		Reason:     fmt.Sprintf("规则 %s 默认行为", rule.Name),
	}
}

// 默认行为
func (crp *CustomRulePolicy) defaultBehavior(cached *CachedAlert, alert *StandardAlert, config *DeduplicationConfig) *DeduplicationResult {
	// 使用严格策略作为默认行为
	strictPolicy := &StrictPolicy{}
	return strictPolicy.ShouldSend(cached, alert, config)
}

// 策略管理器
type PolicyManager struct {
	policies map[string]DeduplicationPolicy
	rules    []*DeduplicationRule
}

// 创建策略管理器
func NewPolicyManager() *PolicyManager {
	pm := &PolicyManager{
		policies: make(map[string]DeduplicationPolicy),
		rules:    make([]*DeduplicationRule, 0),
	}

	// 注册内置策略
	pm.RegisterPolicy(&StrictPolicy{})
	pm.RegisterPolicy(&LoosePolicy{})
	pm.RegisterPolicy(&SeverityBasedPolicy{})

	return pm
}

// 注册策略
func (pm *PolicyManager) RegisterPolicy(policy DeduplicationPolicy) {
	pm.policies[policy.GetName()] = policy
}

// 获取策略
func (pm *PolicyManager) GetPolicy(name string) (DeduplicationPolicy, bool) {
	policy, exists := pm.policies[name]
	return policy, exists
}

// 获取所有策略
func (pm *PolicyManager) GetAllPolicies() map[string]DeduplicationPolicy {
	return pm.policies
}

// 添加自定义规则
func (pm *PolicyManager) AddRule(rule *DeduplicationRule) {
	pm.rules = append(pm.rules, rule)
	
	// 重新创建自定义规则策略
	customPolicy := NewCustomRulePolicy(pm.rules)
	pm.RegisterPolicy(customPolicy)
}

// 删除规则
func (pm *PolicyManager) RemoveRule(name string) bool {
	for i, rule := range pm.rules {
		if rule.Name == name {
			pm.rules = append(pm.rules[:i], pm.rules[i+1:]...)
			
			// 重新创建自定义规则策略
			customPolicy := NewCustomRulePolicy(pm.rules)
			pm.RegisterPolicy(customPolicy)
			
			return true
		}
	}
	return false
}

// 获取所有规则
func (pm *PolicyManager) GetAllRules() []*DeduplicationRule {
	return pm.rules
}

// 验证规则
func (pm *PolicyManager) ValidateRule(rule *DeduplicationRule) error {
	if rule.Name == "" {
		return fmt.Errorf("规则名称不能为空")
	}

	if len(rule.Conditions) == 0 && len(rule.Labels) == 0 {
		return fmt.Errorf("规则必须包含至少一个条件或标签匹配")
	}

	if len(rule.Actions) == 0 {
		return fmt.Errorf("规则必须包含至少一个动作")
	}

	// 验证条件
	for _, condition := range rule.Conditions {
		if condition.Field == "" {
			return fmt.Errorf("条件字段不能为空")
		}
		if condition.Operator == "" {
			return fmt.Errorf("条件操作符不能为空")
		}
		if condition.Operator == "regex" {
			if _, err := regexp.Compile(condition.Value); err != nil {
				return fmt.Errorf("正则表达式无效: %s", condition.Value)
			}
		}
	}

	// 验证动作
	for _, action := range rule.Actions {
		if action.Type == "" {
			return fmt.Errorf("动作类型不能为空")
		}
		if action.Type == "suppress" && action.Duration <= 0 {
			return fmt.Errorf("抑制动作必须指定有效的持续时间")
		}
		if action.Type == "modify_count" && action.MaxCount <= 0 {
			return fmt.Errorf("修改计数动作必须指定有效的最大次数")
		}
	}

	return nil
}

// 全局策略管理器
var GlobalPolicyManager *PolicyManager

// 初始化策略管理器
func InitPolicyManager() {
	GlobalPolicyManager = NewPolicyManager()
}

// 获取全局策略管理器
func GetGlobalPolicyManager() *PolicyManager {
	if GlobalPolicyManager == nil {
		InitPolicyManager()
	}
	return GlobalPolicyManager
}