package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"


}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"


	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1 << 30 - 1
		}
		this.SetSession("uname", uname)
		this.SetSession("pwd", pwd)
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
}


func (this *LoginController) Exit()  {
	this.Ctx.SetCookie("uname", "",-1,"/" )

	this.DelSession("uname")
	this.DelSession("pwd")

	this.Ctx.SetCookie("pwd", "",-1,"/" )
	this.Data["IsLogin"] = false
	this.Redirect("/", 301)
	return

}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, e := ctx.Request.Cookie("pwd")
	if e != nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd
}
