package routers

import (
	"github.com/rh01/beego-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add",&controllers.ReplyController{}, "post:Add")
	beego.Include(&controllers.TopicController{})
	//beego.Router("/topic/view/:id", &controllers.TopicController{}, "get:View")
	// 自动路由
	beego.AutoRouter(&controllers.LoginController{})
	beego.AutoRouter(&controllers.CategoryController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.AutoRouter(&controllers.ReplyController{})
}
