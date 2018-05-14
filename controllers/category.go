package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
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

		error := models.DeleteCategory(id)
		if error!=nil {
			beego.Error(error)
		}

		this.Redirect("/category", 301)
		return

	}

	var e error
	this.Data["Categories"],e = models.GetAllCategories()
	if e != nil {
		beego.Error(e)
	}

}

func (this *CategoryController) Post() {

}
