package main

import (
	_ "github.com/rh01/beego-blog/routers"
	"github.com/astaxie/beego"
	"github.com/rh01/beego-blog/models"
	"github.com/astaxie/beego/orm"
)

func init(){
	models.RegisterDB()
}

func main() {
	// 设置session
	beego.BConfig.WebConfig.Session.SessionOn = true

	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
