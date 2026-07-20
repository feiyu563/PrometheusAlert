package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
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

func AddRecord(source, channel, status, result, summary, originalContent string) error {
	o := orm.NewOrm()
	alertRecord := &AlertRecord{
		Alertname:   source,          // 消息来源 (Prometheus, Zabbix, GitLab等)
		AlertLevel:  channel,         // 转发渠道 (wx, dd, fs, email, alydh等)
		Labels:      status,          // 状态 ("success", "failed")
		Instance:    result,          // 详细结果 (返回的response或Error信息)
		StartsAt:    time.Now().Format("2006-01-02 15:04:05"),
		EndsAt:      time.Now().Format("2006-01-02 15:04:05"),
		Summary:     summary,         // 告警摘要
		Description: originalContent, // 原始消息内容 (RequestBody)
		AlertStatus: status,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	_, err := o.Insert(alertRecord)
	return err
}

type DashboardStats struct {
	TodayReceived    int
	MonthReceived    int
	TodaySentSuccess int
	MonthSentSuccess int
	SourceNames      []string
	SourceCounts     []int
	ChannelNames     []string
	ChannelCounts    []int
	ChannelRows      []StatRow
	SourceRows       []StatRow
}

type StatRow struct {
	Name        string
	Total       int
	Success     int
	Failed      int
	SuccessRate float64
}

func GetChannelName(key string) string {
	switch key {
	case "dd": return "钉钉"
	case "wx": return "企业微信"
	case "workwechat": return "企业微信应用"
	case "fs": return "飞书"
	case "webhook": return "WebHook"
	case "txdx": return "腾讯云短信"
	case "txdh": return "腾讯云电话"
	case "alydx": return "阿里云短信"
	case "alydh": return "阿里云电话"
	case "bddx": return "百度云短信"
	case "rlydh": return "容联云电话"
	case "7moordx": return "七陌短信"
	case "7moordh": return "七陌语音电话"
	case "email": return "Email"
	case "tg": return "Telegram"
	case "rl": return "百度Hi(如流)"
	case "bark": return "Bark(iPhone推送)"
	case "voice": return "语音播报"
	case "fsapp": return "飞书自建应用"
	case "kafka": return "Kafka"
	}
	return key
}

func GetDashboardStats() (DashboardStats, error) {
	o := orm.NewOrm()
	var stats DashboardStats

	stats.SourceNames = make([]string, 0)
	stats.SourceCounts = make([]int, 0)
	stats.ChannelNames = make([]string, 0)
	stats.ChannelCounts = make([]int, 0)
	stats.ChannelRows = make([]StatRow, 0)
	stats.SourceRows = make([]StatRow, 0)

	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// 1. Today's stats
	var todayRes []orm.Params
	_, err := o.Raw("SELECT alert_status, COUNT(*) as c FROM alert_record WHERE created_time >= ? GROUP BY alert_status", todayStart).Values(&todayRes)
	if err == nil {
		for _, r := range todayRes {
			status := r["alert_status"].(string)
			countStr := r["c"].(string)
			count, _ := strconv.Atoi(countStr)
			stats.TodayReceived += count
			if status == "success" || status == "resolved" || status == "firing" {
				stats.TodaySentSuccess += count
			}
		}
	}

	// 2. Month's stats
	var monthRes []orm.Params
	_, err = o.Raw("SELECT alert_status, COUNT(*) as c FROM alert_record WHERE created_time >= ? GROUP BY alert_status", monthStart).Values(&monthRes)
	if err == nil {
		for _, r := range monthRes {
			status := r["alert_status"].(string)
			countStr := r["c"].(string)
			count, _ := strconv.Atoi(countStr)
			stats.MonthReceived += count
			if status == "success" || status == "resolved" || status == "firing" {
				stats.MonthSentSuccess += count
			}
		}
	}

	// 3. Historical Source Stats (Group by alertname & alert_status)
	var sourceRes []orm.Params
	_, err = o.Raw("SELECT alertname, alert_status, COUNT(*) as c FROM alert_record GROUP BY alertname, alert_status").Values(&sourceRes)
	if err == nil {
		sourceMap := make(map[string]*StatRow)
		for _, r := range sourceRes {
			name := r["alertname"].(string)
			if name == "" {
				name = "Prometheus"
			}
			status := r["alert_status"].(string)
			countStr := r["c"].(string)
			count, _ := strconv.Atoi(countStr)

			row, ok := sourceMap[name]
			if !ok {
				row = &StatRow{Name: name}
				sourceMap[name] = row
			}
			row.Total += count
			if status == "success" || status == "resolved" || status == "firing" {
				row.Success += count
			} else {
				row.Failed += count
			}
		}

		for _, row := range sourceMap {
			if row.Total > 0 {
				row.SuccessRate = float64(row.Success) / float64(row.Total) * 100
			}
			stats.SourceRows = append(stats.SourceRows, *row)
			stats.SourceNames = append(stats.SourceNames, row.Name)
			stats.SourceCounts = append(stats.SourceCounts, row.Total)
		}
	}

	// 4. Historical Channel Stats (Group by alert_level & alert_status)
	var channelRes []orm.Params
	_, err = o.Raw("SELECT alert_level, alert_status, COUNT(*) as c FROM alert_record GROUP BY alert_level, alert_status").Values(&channelRes)
	if err == nil {
		channelMap := make(map[string]*StatRow)
		for _, r := range channelRes {
			name := r["alert_level"].(string)
			if name == "" {
				continue
			}
			status := r["alert_status"].(string)
			countStr := r["c"].(string)
			count, _ := strconv.Atoi(countStr)

			row, ok := channelMap[name]
			if !ok {
				row = &StatRow{Name: name}
				channelMap[name] = row
			}
			row.Total += count
			if status == "success" || status == "resolved" || status == "firing" {
				row.Success += count
			} else {
				row.Failed += count
			}
		}

		for _, row := range channelMap {
			if row.Total > 0 {
				row.SuccessRate = float64(row.Success) / float64(row.Total) * 100
			}
			// Mapped name for display
			mappedRow := *row
			mappedRow.Name = GetChannelName(row.Name)
			stats.ChannelRows = append(stats.ChannelRows, mappedRow)
			
			stats.ChannelNames = append(stats.ChannelNames, mappedRow.Name)
			stats.ChannelCounts = append(stats.ChannelCounts, row.Total)
		}
	}

	return stats, nil
}

type DisplayRecord struct {
	Time            string
	Source          string
	Channel         string
	Status          string
	Summary         string
	OriginalPayload string
	Result          string
}

func (r *AlertRecord) ToDisplay() DisplayRecord {
	isOldRecord := r.Alertname != "Prometheus" && r.Alertname != "Zabbix" && r.Alertname != "GitLab"

	if isOldRecord {
		return DisplayRecord{
			Time:            r.CreatedTime.Format("2006-01-02 15:04:05"),
			Source:          "Prometheus (" + r.Alertname + ")",
			Channel:         "默认或路由分发",
			Status:          r.AlertStatus, // firing / resolved
			Summary:         r.Summary,
			OriginalPayload: r.Labels, // 原 Prometheus Labels Json
			Result:          "目标主机: " + r.Instance,
		}
	}

	return DisplayRecord{
		Time:            r.CreatedTime.Format("2006-01-02 15:04:05"),
		Source:          r.Alertname,
		Channel:         GetChannelName(r.AlertLevel),
		Status:          r.AlertStatus, // success / failed
		Summary:         r.Summary,
		OriginalPayload: r.Description, // 原始 RequestBody
		Result:          r.Instance,    // 详细结果
	}
}

func GetRecordPage(start, length int, searchVal string) ([]*AlertRecord, int64, int64, error) {
	o := orm.NewOrm()
	Record_all := make([]*AlertRecord, 0)
	qs := o.QueryTable("AlertRecord")

	// Get total count
	total, err := qs.Count()
	if err != nil {
		return nil, 0, 0, err
	}

	// Filter if searchVal is present
	if searchVal != "" {
		cond := orm.NewCondition()
		cond = cond.Or("Alertname__icontains", searchVal).
			Or("AlertLevel__icontains", searchVal).
			Or("Labels__icontains", searchVal).
			Or("Instance__icontains", searchVal).
			Or("Summary__icontains", searchVal).
			Or("Description__icontains", searchVal).
			Or("AlertStatus__icontains", searchVal)
		qs = qs.SetCond(cond)
	}

	// Get filtered count
	filtered, err := qs.Count()
	if err != nil {
		return nil, 0, 0, err
	}

	// Fetch page
	_, err = qs.OrderBy("-id").Offset(start).Limit(length).All(&Record_all)
	return Record_all, total, filtered, err
}
