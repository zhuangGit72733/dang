package controllers

import (
	"dang/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type AdvicesController struct {
	beego.Controller
}
type Message struct {
	Status   bool
	CreateId int64
}

func (c *AdvicesController) Advice() {
	var v models.Advices
	v.Name = c.Input().Get("name")
	v.Email = c.Input().Get("email")
	phone := c.Input().Get("phone")
	v.Phone, _ = strconv.Atoi(phone)
	v.Unit = c.Input().Get("unit")
	v.Content = c.Input().Get("content")
	v.CreatedAt = time.Now().String()
	v.UpdatedAt = time.Now().String()
	if id, err := models.PostAdvice(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		logs.Debug(v)
		message := &Message{Status: true, CreateId: id}
		c.Data["json"] = message
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
