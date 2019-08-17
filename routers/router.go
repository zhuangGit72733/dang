package routers

import (
	"dang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("api/videos", &controllers.VideosController{}, "Get:Video")
	beego.Router("api/advices/create", &controllers.AdvicesController{}, "Post:Advice")
	beego.Router("api/members", &controllers.MembersController{}, "Get:Sum")
	beego.Router("api/members/create", &controllers.MembersController{}, "Post:CreateMember")
	beego.Router("api/hongbao", &controllers.MapsController{}, "Get:San")
	beego.Router("api/hongcun", &controllers.MapsController{}, "Get:Cun")
	beego.Router("api/fengcai", &controllers.MiensController{}, "Get:Fengcai")
	beego.Router("api/gongzuoshi", &controllers.WorkersController{}, "Get:Gongzuoshi")
	beego.Router("api/branches", &controllers.BranchesController{}, "Get:Branches")
}
