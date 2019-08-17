package controllers

import (
	"dang/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type WorkersController struct {
	beego.Controller
}

func (c *WorkersController) Gongzuoshi() {
	works, workers, err := models.Gongzuoshi()
	if err != nil {
		logs.Debug("workers controller err:%v", err)
	}
	for k, v := range works {
		var Arr []orm.Params
		for _, val := range workers {
			if v["id"] == val["work_id"] {
				Arr = append(Arr, val)
			}
		}
		works[k]["message"] = Arr
	}
	c.Data["json"] = works
	c.ServeJSON()
}
