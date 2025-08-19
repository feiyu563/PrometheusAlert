package controllers

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func LogsSign() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// 转换时间戳到时间字符串
func GetTime(timeStr interface{}, timeFormat ...string) string {
	var R_Time string
	//判断传入的timeStr是否为float64类型，如gerrit消息中时间戳就是float64
	switch timeStr.(type) {
	case string:
		S_Time, _ := strconv.ParseInt(timeStr.(string), 10, 64)
		if len(timeFormat) == 0 {
			timeFormat = append(timeFormat, "2006-01-02T15:04:05")
		}
		if len(timeStr.(string)) == 13 {
			R_Time = time.Unix(S_Time/1000, 0).Format(timeFormat[0])
		} else {
			R_Time = time.Unix(S_Time, 0).Format(timeFormat[0])
		}
	case float64:
		if len(timeFormat) == 0 {
			timeFormat = append(timeFormat, "2006-01-02T15:04:05")
		}
		R_Time = time.Unix(int64(timeStr.(float64)), 0).Format(timeFormat[0])
	}
	return R_Time
}

// 转换时间为持续时长
func GetTimeDuration(startTime string,endTime string) string {
	var tm = "N/A"
	if startTime != "" && endTime != "" {
		starT1 := startTime[0:10]
		starT2 := startTime[11:19]
		starT3 := starT1 + " " + starT2
		startm2, err := time.Parse("2006-01-02 15:04:05", starT3)
		if err != nil {
			return tm // 如果解析失败，则返回N/A
		}

		endT1 := endTime[0:10]
		endT2 := endTime[11:19]
		endT3 := endT1 + " " + endT2
		endm2, err := time.Parse("2006-01-02 15:04:05", endT3)
		if err != nil {
			return tm // 如果解析失败，则返回N/A
		}

		sub := endm2.UTC().Sub(startm2.UTC())

		t := int64(sub.Seconds())
		if t >= 86400 {
			days := t / 86400
			hours := (t % 86400) / 3600
			tm = fmt.Sprintf("%dd%dh", days, hours)
		} else {
			hours := t / 3600
			minutes := (t % 3600) / 60
			if hours > 0 {
				tm = fmt.Sprintf("%dh%dm", hours, minutes)
			} else {
				// 如果小时为0，则只显示分钟和秒
				seconds := t % 60
				tm = fmt.Sprintf("%dm%ds", minutes, seconds)
				if minutes == 0 {
					// 如果分钟也为0，则只显示秒
					tm = fmt.Sprintf("%ds", seconds)
				}
			}
		}
	}
	return tm
}

// 转换任意时区到CST
func GetCSTtime(date string) string {
	var t time.Time
	if date == "" {
		// 获取当前时间并转换为 CST
		t = time.Now().In(cstLoc)
	} else {
		parsedTime, err := parseDate(date)
		if err != nil {
			// 处理错误，例如返回空字符串或日志记录
			return ""
		}
		// 转换为 CST 时区
		t = parsedTime.In(cstLoc)
	}
	return t.Format("2006-01-02 15:04:05")
}

// 解析日期字符串，支持带时区和固定格式
func parseDate(date string) (time.Time, error) {
	// 尝试常见带时区的格式（如 RFC3339）
	t, err := time.Parse(time.RFC3339, date)
	if err == nil {
		return t, nil
	}
	// 尝试无时区格式，假设为 UTC
	t, err = time.ParseInLocation("2006-01-02 15:04:05", date, time.UTC)
	if err == nil {
		return t, nil
	}
	// 可根据需要添加更多格式
	return time.Time{}, fmt.Errorf("无法解析时间：%s", date)
}

func TimeFormat(timestr, format string) string {
	returnTime, err := time.Parse("2006-01-02T15:04:05.999999999Z", timestr)
	if err != nil {
		returnTime, err = time.Parse("2006-01-02T15:04:05.999999999+08:00", timestr)
	}
	if err != nil {
		return err.Error()
	} else {
		return returnTime.Format(format)
	}
}

// 获取用户号码
func GetUserPhone(neednum int) string {
	//判断是否存在user.csv文件
	Num := beego.AppConfig.String("defaultphone")
	Today := time.Now()
	//判断当前时间是否大于10点,大于10点取当天值班号码,小于10点取前一天值班号码
	DayString := ""
	if time.Now().Hour() >= 10 {
		//取当天值班号码
		DayString = Today.Format("2006年1月2日")
	} else {
		//取前一天值班号码
		DayString = Today.AddDate(0, 0, -1).Format("2006年1月2日")
	}
	_, err := os.Stat("user.csv")
	if err == nil {
		f, err := os.Open("user.csv")
		if err != nil {
			logs.Error(err.Error())
		}
		defer f.Close()
		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
			if err != nil {
				if err.Error() != "EOF" {
					logs.Error(err.Error())
				}
				break
			}
			if strings.Contains(line, DayString) {
				x := strings.Split(line, ",")
				Num = x[neednum]
				break
			}
		}
		f.Close()
	} else {
		logs.Error(err.Error())
	}
	return Num
}

// 随机返回
func DoBalance(instances []string) string {
	if len(instances) == 0 {
		logs.Error("no instances for rand")
		return ""
	}
	lens := len(instances)
	index := rand.Intn(lens)
	inst := instances[index]
	return inst
}
