package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Branches struct {
	Id   int64
	Name string `orm:"size(128)"`
}
type Activities struct {
	Id       int64
	BranchId string
	Title    string `orm:"size(128)"`
	Image    string
}

func init() {
	orm.RegisterModel(new(Branches), new(Activities))
}
func Braches() (maps []orm.Params, mapers []orm.Params, err error) {

	maps, mapers, err = getActivity("select id,name from branches")
	return maps, mapers, err
}
func getActivity(query string) (maps []orm.Params, mapers []orm.Params, err error) {
	o := orm.NewOrm()
	num, err := o.Raw(query).Values(&maps)
	if err != nil || num < 1 {
		return nil, nil, err
	}
	query = fmt.Sprintf("select branch_id,title,image from activities")
	num, err = o.Raw(query).Values(&mapers)
	if err != nil || num < 1 {
		return nil, nil, err
	}
	return maps, mapers, err
}
