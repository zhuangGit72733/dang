package controllers

import (
	"dang/config"
	"dang/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type VideosController struct {
	beego.Controller
}

func (c *VideosController) Video() {
	url, _ := config.InitConfig()
	data, err := models.GetVideos()
	if err != nil {
		logs.Debug("videos controller err:%v", err)
	}
	for k, v := range data {
		data[k]["video"] = fmt.Sprintf("%v/uploads/videos/%v", url, v["video"])
		if data[k]["img"] != nil {
			data[k]["img"] = fmt.Sprintf("%v/uploads/activity/%v", url, v["img"])
		}
	}
	c.Data["json"] = data
	c.ServeJSON()
}
