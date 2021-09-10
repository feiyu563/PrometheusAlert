package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	//判断是否为退出操作
	if c.Input().Get("exit") == "true" {
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
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
	if beego.AppConfig.String("login_user") == username && beego.AppConfig.String("login_password") == password {
		maxage := 0
		if autologin {
			maxage = 1<<31 - 1
		}
		c.Ctx.SetCookie("username", username, maxage, "/")
		c.Ctx.SetCookie("password", password, maxage, "/")
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

//检查cookie是否为登录状态
func checkAccount(mycookie *context.Context) bool {
	ck, err := mycookie.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value
	ck, err = mycookie.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value
	return beego.AppConfig.String("login_user") == username && beego.AppConfig.String("login_password") == password
}
