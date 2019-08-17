package controllers

import (
	"dang/config"
	"dang/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type BranchesController struct {
	beego.Controller
}

func (c *BranchesController) Branches() {
	branches, activities, err := models.Braches()
	if err != nil {
		logs.Debug("Miens controller err:%v", err)
	}
	for k, v := range branches {
		var Arr []orm.Params
		for _, val := range activities {
			if v["id"] == val["branch_id"] {
				ImagePath(val)
				Arr = append(Arr, val)
			}
		}
		branches[k]["message"] = Arr
	}
	c.Data["json"] = branches
	c.ServeJSON()
}
func ImagePath(start orm.Params) {
	url, _ := config.InitConfig()
	start["image"] = fmt.Sprintf("%v/uploads/activity/%v", url, start["image"])
}
