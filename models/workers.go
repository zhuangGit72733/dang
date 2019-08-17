package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Works struct {
	Id   int64
	name string
}
type Workers struct {
	Id      int64
	WorkId  string
	Title   string `orm:"size(128)"`
	Content string
}

func init() {
	orm.RegisterModel(new(Works), new(Workers))
}
func Gongzuoshi() (maps []orm.Params, mapers []orm.Params, err error) {
	maps, mapers, err = getWork("select id,name from works")
	return maps, mapers, err
}
func getWork(query string) (maps []orm.Params, mapers []orm.Params, err error) {
	o := orm.NewOrm()
	num, err := o.Raw(query).Values(&maps)
	if err != nil || num < 1 {
		return nil, nil, err
	}
	query = fmt.Sprintf("select work_id,title,content from workers")
	num, err = o.Raw(query).Values(&mapers)
	if err != nil || num < 1 {
		return nil, nil, err
	}
	return maps, mapers, err
}
