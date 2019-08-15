package main

import (
	"dang/config"
	_ "dang/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/tangqiao?charset=utf8")
	addr, err := config.InitConfig()
	if err != nil {
		logs.Debug("Init config addr fail,err:%v", err)
	}
	fmt.Println(addr)
	beego.Run()
}
