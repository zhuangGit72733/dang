package controllers

import (
	"dang/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MapsController struct {
	beego.Controller
}

func (c *MapsController) San() {
	maps, mapers, err := models.GetSan()
	if err != nil {
		logs.Debug("mapers controller err:%v", err)
	}
	for k, v := range maps {
		var Arr []orm.Params
		for _, val := range mapers {
			if v["id"] == val["map_id"] {
				Arr = append(Arr, val)
			}
		}
		maps[k]["message"] = Arr
	}
	c.Data["json"] = maps
	c.ServeJSON()
}
func (c *MapsController) Cun() {
	maps, mapers, err := models.GetCun()
	if err != nil {
		logs.Debug("mapers controller err:%v", err)
	}
	for k, v := range maps {
		var Arr []orm.Params
		for _, val := range mapers {
			if v["id"] == val["map_id"] {
				Arr = append(Arr, val)
			}
		}
		maps[k]["message"] = Arr
	}
	c.Data["json"] = maps
	c.ServeJSON()
}
