package controllers

import (
	"io"
	"os"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// ServiceLog renders the page views/servicelog.html
func (c *MainController) ServiceLog() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "servicelog.html"
}

type LogResponse struct {
	Content   string `json:"content"`
	Truncated bool   `json:"truncated"`
	Size      int64  `json:"size"`
}

// ServiceLogData reads logs/prometheusalertcenter.log and returns JSON
func (c *MainController) ServiceLogData() {
	if !CheckAccount(c.Ctx) {
		c.Data["json"] = "unauthorized"
		c.ServeJSON()
		return
	}

	logpath := beego.AppConfig.String("logpath")
	if logpath == "" {
		logpath = "logs/prometheusalertcenter.log"
	}

	file, err := os.Open(logpath)
	if err != nil {
		logs.Error("Failed to open log file:", err.Error())
		c.Data["json"] = LogResponse{
			Content:   "无法打开日志文件: " + err.Error() + "\n请检查配置文件 app.conf 中的 logpath 是否正确设置，或者是否有该文件的读取权限。",
			Truncated: false,
			Size:      0,
		}
		c.ServeJSON()
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		c.Data["json"] = LogResponse{
			Content:   "无法获取日志文件大小: " + err.Error(),
			Truncated: false,
			Size:      0,
		}
		c.ServeJSON()
		return
	}

	fileSize := stat.Size()
	threshold := int64(1024 * 1024) // 1MB 安全阈值
	truncated := false

	var reader io.Reader = file
	if fileSize > threshold {
		truncated = true
		_, err = file.Seek(-threshold, io.SeekEnd) // 强行移动指针至距离末尾 1MB 处
		if err != nil {
			logs.Error("Failed to seek log file:", err.Error())
			c.Data["json"] = LogResponse{
				Content:   "定位日志文件失败: " + err.Error(),
				Truncated: false,
				Size:      fileSize,
			}
			c.ServeJSON()
			return
		}
	}

	buf := make([]byte, threshold)
	n, err := reader.Read(buf)
	if err != nil && err != io.EOF {
		c.Data["json"] = LogResponse{
			Content:   "读取日志文件失败: " + err.Error(),
			Truncated: false,
			Size:      fileSize,
		}
		c.ServeJSON()
		return
	}

	// 统一转换字符串返回
	logContent := string(buf[:n])

	c.Data["json"] = LogResponse{
		Content:   logContent,
		Truncated: truncated,
		Size:      fileSize,
	}
	c.ServeJSON()
}
