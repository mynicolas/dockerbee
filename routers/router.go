package routers

import (
	"dockerbee/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/polymer", "polymer")
	beego.SetStaticPath("/js", "js")
	beego.SetStaticPath("/css", "css")
	beego.SetStaticPath("/img", "img")

	beego.Router("/test", &controllers.TestController{})
  beego.Router("/", &controllers.MainController{})
  beego.Router("/overview", &controllers.OverviewController{})
  beego.Router("/containers", &controllers.ContainersController{})
  beego.Router("/images", &controllers.ImagesController{})

  beego.Router("/actions/pause", &controllers.ContainersController{}, "post:Pause")
  beego.Router("/actions/unpause", &controllers.ContainersController{}, "post:Unpause")
  beego.Router("/actions/stop", &controllers.ContainersController{}, "post:Stop")
  beego.Router("/actions/start", &controllers.ContainersController{}, "post:Start")
  beego.Router("/actions/restart", &controllers.ContainersController{}, "post:Restart")
  beego.Router("/actions/kill", &controllers.ContainersController{}, "post:Kill")
  beego.Router("/actions/remove", &controllers.ContainersController{}, "post:Remove")

}
