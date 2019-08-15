package routers

import (
	"dang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/videos", &controllers.VideosController{}, "Get:Video")
}
