package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Miens struct {
	Id        int64
	name      string
	CreatedAt string
	UpdatedAt string
}
type Parties struct {
	Id        int64
	MienId    string
	Title     string `orm:"size(128)"`
	Content   string
	Img       string
	CreatedAt string
	UpdatedAt string
}

func init() {
	orm.RegisterModel(new(Miens), new(Parties))
}
func Fengcai() (maps []orm.Params, mapers []orm.Params, err error) {
	maps, mapers, err = getMien("select id,name from miens")
	return maps, mapers, err
}
func getMien(query string) (maps []orm.Params, mapers []orm.Params, err error) {
	o := orm.NewOrm()
	num, err := o.Raw(query).Values(&maps)
	if err != nil || num < 1 {
		return nil, nil, err
	}
	query = fmt.Sprintf("select mien_id,img,title,content from parties")
	num, err = o.Raw(query).Values(&mapers)
	if err != nil || num < 1 {
		return nil, nil, err
	}
	return maps, mapers, err
}
