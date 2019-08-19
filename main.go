package main

import (
	_ "dang/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//注册数据库到ORM
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/tangqiao?charset=utf8mb4")
	//跨域解决
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true}))
	beego.Run()
}
