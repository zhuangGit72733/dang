package controllers

import (
	"dang/config"
	"dang/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type VideosController struct {
	beego.Controller
}

func (c *VideosController) Video() {
	data, err := models.GetVideos()
	if err != nil {
		logs.Debug("videos controller err:%v", err)
	}
	c.Data["json"] = addPath(data)
	c.ServeJSON()
}
func addPath(data []orm.Params) []orm.Params {
	url, _ := config.InitConfig()
	for k, v := range data {
		data[k]["video"] = fmt.Sprintf("%v/uploads/videos/%v", url, v["video"])
		if data[k]["img"] != nil {
			data[k]["img"] = fmt.Sprintf("%v/uploads/activity/%v", url, v["img"])
		}
	}
	return data
}
