package controllers

import (
	"github.com/astaxie/beego"
	"github.com/rh01/beego-blog/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	var err error
	this.Data["Topics"], err = models.GetAllTopics(this.Input().Get("cate"), false)
	this.Data["Categories"],err = models.GetAllCategories()

	if err != nil{
		beego.Error(err)
	}

	this.TplName = "index.html"

}
