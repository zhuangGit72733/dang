package controllers

import (
	"dang/config"
	"dang/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strings"
)

type ArticlesController struct {
	beego.Controller
}

func getArticles(type_id string) (Arr []orm.Params) {
	articles, err := models.Articles()
	url, _ := config.InitConfig()
	if err != nil {
		logs.Debug("Articles controller err:%v", err)
	}
	for _, val := range articles {
		if val["type_id"] == type_id {
			val["logo"] = fmt.Sprintf("%v/uploads/activity/%v", url, val["logo"])
			val["er"] = fmt.Sprintf("%v/uploads/activity/%v", url, val["er"])
			if val["img"] != nil {
				val["img"] = ImageArrPath(val["img"])
			}
			Arr = append(Arr, val)
		}
	}
	return
}
func ImageArrPath(inter interface{}) (arr []string) {
	var str string
	url, _ := config.InitConfig()
	str = inter.(string)
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)
	str = strings.Replace(str, "\"", "", -1)
	arr = strings.Split(str, ",")
	for k, _ := range arr {
		arr[k] = fmt.Sprintf("%v/uploads/activity/%v", url, arr[k])
	}
	return
}
func (c *ArticlesController) Dangjian() {
	c.Data["json"] = getArticles("1")
	c.ServeJSON()
}
func (c *ArticlesController) Wenming() {
	c.Data["json"] = getArticles("5")
	c.ServeJSON()
}

func (c *ArticlesController) Zhijian() {
	c.Data["json"] = getArticles("10")
	c.ServeJSON()
}
func (c *ArticlesController) Shici() {
	c.Data["json"] = getArticles("16")
	c.ServeJSON()
}
func (c *ArticlesController) Huodong() {
	c.Data["json"] = getArticles("9")
	c.ServeJSON()
}
func (c *ArticlesController) San() {
	types, _ := models.Types()
	articles, err := models.Articles()
	if err != nil {
		logs.Debug("Articles controller err:%v", err)
	}
	for key, v := range types {
		var Arr []orm.Params
		for _, val := range articles {
			if val["type_id"] == v["id"] {
				Arr = append(Arr, val)
			}
		}
		types[key]["message"] = Arr
	}

	c.Data["json"] = types
	c.ServeJSON()
}
