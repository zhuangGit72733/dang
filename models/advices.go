package models

import (
	"github.com/astaxie/beego/orm"
)

type Advices struct {
	Id        int64
	Name      string `orm:"size(128)"`
	Email     string `orm:"size(128)"`
	Phone     int
	Unit      string
	Content   string
	CreatedAt string
	UpdatedAt string
}

func init() {
	orm.RegisterModel(new(Advices))
}

func PostAdvice(m *Advices) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
