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

  beego.Router("/actions", &controllers.ContainersController{}, "post:Action")

}
