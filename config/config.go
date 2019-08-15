package config

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	myConfig = &MyConfig{}
)

type MyConfig struct {
	HttpUrl string
}

func InitConfig() (addr string, err error) {
	addr = beego.AppConfig.String("http_url")
	if len(addr) == 0 {
		logs.Debug("invalid conf http_url")
	}
	myConfig.HttpUrl = addr
	return
}
