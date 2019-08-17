package controllers

import (
	"dang/config"
	"dang/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MiensController struct {
	beego.Controller
}

func (c *MiensController) Fengcai() {
	mien, parties, err := models.Fengcai()
	if err != nil {
		logs.Debug("Miens controller err:%v", err)
	}
	for k, v := range mien {
		var Arr []orm.Params
		for _, val := range parties {
			if v["id"] == val["mien_id"] {
				ImgPath(val)
				Arr = append(Arr, val)
			}
		}
		mien[k]["message"] = Arr
	}
	c.Data["json"] = mien
	c.ServeJSON()
}
func ImgPath(start orm.Params) {
	url, _ := config.InitConfig()
	start["img"] = fmt.Sprintf("%v/uploads/activity/%v", url, start["img"])
}
