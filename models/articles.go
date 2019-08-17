package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Type struct {
	Id   int64
	name string
}
type Article struct {
	Id        int64
	TypeId    string
	Title     string `orm:"size(128)"`
	Content   string
	Img       string
	Logo      string
	er        string
	small     string
	CreatedAt string
	UpdatedAt string
}

func init() {
	orm.RegisterModel(new(Type), new(Article))
}
func Articles() (mapers []orm.Params, err error) {
	mapers, err = getArticle()
	return mapers, err
}
func Types() (mapers []orm.Params, err error) {
	o := orm.NewOrm()
	query := fmt.Sprintf("select id,name from types where id>11 and id<15")
	num, err := o.Raw(query).Values(&mapers)
	if err != nil || num < 1 {
		return nil, err
	}
	logs.Debug(num)
	return mapers, err
}
func getArticle() (mapers []orm.Params, err error) {
	o := orm.NewOrm()
	query := fmt.Sprintf("select * from articles order by updated_at desc")
	num, err := o.Raw(query).Values(&mapers)
	if err != nil || num < 1 {
		return nil, err
	}
	logs.Debug(num)
	return mapers, err
}
