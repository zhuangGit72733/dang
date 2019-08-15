package routers

import (
	"dang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("api/videos", &controllers.VideosController{}, "Get:Video")
	beego.Router("api/advices/create", &controllers.AdvicesController{}, "Post:Advice")
}
