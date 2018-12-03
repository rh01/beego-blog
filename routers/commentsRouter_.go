package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beeblog/controllers:TopicController"] = append(beego.GlobalControllerRouter["beeblog/controllers:TopicController"],
		beego.ControllerComments{
			Method: "View",
			Router: `/topic/view/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
