package controllers

import (
	"dang/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type MembersController struct {
	beego.Controller
}
type Sum struct {
	Sum int64
}

func (c *MembersController) Sum() {
	sum := models.GetSum()
	c.Data["json"] = &Sum{Sum: sum}
	c.ServeJSON()
}
func (c *MembersController) CreateMember() {
	var v models.Members
	v.Name = c.Input().Get("name")
	v.Leader = c.Input().Get("leader")
	phone := c.Input().Get("phone")
	v.Phone, _ = strconv.Atoi(phone)
	v.Unit = c.Input().Get("unit")
	sum := c.Input().Get("sum")
	v.Sum, _ = strconv.Atoi(sum)
	v.CreatedAt = time.Now().String()
	v.UpdatedAt = time.Now().String()
	if id, err := models.CreateMembers(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		logs.Debug(v)
		message := &Message{Status: true, CreateId: id}
		c.Data["json"] = message
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
