package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/rh01/beego-blog/controllers:TopicController"] = append(beego.GlobalControllerRouter["github.com/rh01/beego-blog/controllers:TopicController"],
        beego.ControllerComments{
            Method: "View",
            Router: `/topic/view/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
