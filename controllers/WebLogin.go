package controllers

import (
	"crypto/sha256"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {
	title := beego.AppConfig.String("title")
	if title == "" {
		title = "PrometheusAlert"
	}
	c.Data["AppTitle"] = title
}

func (c *LoginController) Get() {
	//判断是否为退出操作
	if c.Input().Get("exit") == "true" {
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("logintoken", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/") // 清除旧版本可能残留的明文密码 cookie
		c.Redirect("/", 302)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	//获取用户名和密码信息
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	autologin := c.Input().Get("autologin") == "on"
	//判断用户名密码是否正确
	cfgUser := beego.AppConfig.String("login_user")
	cfgPwd := beego.AppConfig.String("login_password")
	if cfgUser == username && cfgPwd == password {
		maxage := 0
		if autologin {
			maxage = 1<<31 - 1
		}
		token := fmt.Sprintf("%x", sha256.Sum256([]byte(username+password+"PrometheusAlert")))
		c.Ctx.SetCookie("username", username, maxage, "/")
		c.Ctx.SetCookie("logintoken", token, maxage, "/")
		c.Redirect("/", 301)
		return
	} else {
		//flash.Data["error"]=true
		//c.Redirect("/login",301)
		//return
		c.TplName = "login.html"
		c.Data["loginerror"] = true
		//err:=c.Render()
		//if err!=nil {
		//	beego.Error(err)
		//}
	}

	//return
}

// 检查cookie是否为登录状态
func CheckAccount(mycookie *context.Context) bool {
	ckUser, err := mycookie.Request.Cookie("username")
	if err != nil {
		return false
	}
	ckToken, err := mycookie.Request.Cookie("logintoken")
	if err != nil {
		return false
	}
	cfgUser := beego.AppConfig.String("login_user")
	cfgPwd := beego.AppConfig.String("login_password")
	expectedToken := fmt.Sprintf("%x", sha256.Sum256([]byte(cfgUser+cfgPwd+"PrometheusAlert")))
	return ckUser.Value == cfgUser && ckToken.Value == expectedToken
}
