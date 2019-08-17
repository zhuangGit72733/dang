package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Maps struct {
	Id   int64
	Name string `orm:"size(128)"`
}
type Mapers struct {
	Id      int64
	MapId   string
	Title   string `orm:"size(128)"`
	Content string
}

func init() {
	orm.RegisterModel(new(Maps), new(Mapers))
}
func GetSan() (maps []orm.Params, mapers []orm.Params, err error) {

	maps, mapers, err = getComment("select id,name from maps where id <4")
	return maps, mapers, err
}
func GetCun() (maps []orm.Params, mapers []orm.Params, err error) {
	maps, mapers, err = getComment("select id,name from maps where id >=4")
	return maps, mapers, err
}
func getComment(query string) (maps []orm.Params, mapers []orm.Params, err error) {
	o := orm.NewOrm()
	num, err := o.Raw(query).Values(&maps)
	if err != nil || num < 1 {
		return nil, nil, err
	}
	query = fmt.Sprintf("select map_id,title,content from mapers order by  updated_at desc")
	num, err = o.Raw(query).Values(&mapers)
	if err != nil || num < 1 {
		return nil, nil, err
	}
	return maps, mapers, err
}
