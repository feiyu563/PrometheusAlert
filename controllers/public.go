package controllers

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//转换UTC时区到CST
func GetCSTtime(date string) string {
	var tm string
	tm = time.Now().Format("2006-01-02 15:04:05")
	if date != "" {
		T1 := date[0:10]
		T2 := date[11:19]
		T3 := T1 + " " + T2
		tm2, _ := time.Parse("2006-01-02 15:04:05", T3)
		h, _ := time.ParseDuration("-1h")
		tm3 := tm2.Add(-8 * h)
		tm = tm3.Format("2006-01-02 15:04:05")

	}
	return tm
}

func TimeFormat(timestr,format string) string {
	returnTime,err:=time.Parse("2006-01-02T15:04:05.999999999Z",timestr)
	if err!=nil {
		logs.Error("自定义模板时间格式转换失败："+err.Error())
	}
	return returnTime.Format(format)
}

//获取用户号码
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
