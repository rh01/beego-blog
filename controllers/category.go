package controllers

import (
	"github.com/astaxie/beego"
	"github.com/rh01/beego-blog/models"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get()  {

	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	op := this.Input().Get("op")
	switch op {
	case "add":
		if !checkAccount(this.Ctx) {
			this.Redirect("/login", 302)
			return
		}
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return

	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DeleteOneCategory(id)
		if err!=nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return

	}

	var e error
	this.Data["Categories"],e = models.GetAllCategories()
	if e != nil {
		beego.Error(e)
	}

}

func (this *CategoryController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

}

func (this *CategoryController)  Del()  {

	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	cid := this.Input().Get("id")
	e := models.DeleteOneCategory(cid)
	if e != nil {
		beego.Error(e)
	}

	this.Redirect("/category", 302)
	return

}