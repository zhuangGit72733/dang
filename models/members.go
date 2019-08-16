package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Members struct {
	Id        int64
	Name      string `orm:"size(128)"`
	Phone     int
	Unit      string
	Leader    string `orm:"size(128)"`
	Sum       int
	CreatedAt string
	UpdatedAt string
}

func init() {
	orm.RegisterModel(new(Members))
}

func CreateMembers(m *Members) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
func GetSum() (sum int64) {
	var maps []orm.Params
	o := orm.NewOrm()
	query := fmt.Sprintf("select * from members")
	sum, err := o.Raw(query).Values(&maps)
	if err != nil {
		logs.Debug("get member error:%v", err)
	}
	return
}
