package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Videos struct {
	Id        int64
	Title     string `orm:"size(128)"`
	Video     string `orm:"size(128)"`
	CreatedAt string
	UpdatedAt string
}

func init() {
	orm.RegisterModel(new(Videos))
}

func GetVideos() (maps []orm.Params, err error) {
	o := orm.NewOrm()
	query := fmt.Sprintf("select * from videos")
	num, err := o.Raw(query).Values(&maps)
	if err != nil || num < 1 {
		return nil, err
	}
	return maps, err
}
