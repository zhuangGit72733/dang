package controllers

import (
	"dang/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type VideosController struct {
	beego.Controller
}

func (c *VideosController) Video() {
	data, err := models.GetVideos()
	if err != nil {
		logs.Debug("videos controller err:%v", err)
	}
	c.Data["json"] = data
	c.ServeJSON()
}
