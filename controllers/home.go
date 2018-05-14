package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	var err error
	this.Data["Topics"], err = models.GetAllTopics(false)
	if err != nil{
		beego.Error(err)
	}

	this.TplName = "index.html"

}
