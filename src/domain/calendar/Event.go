package calendar

import "github.com/astaxie/beego/orm"

type Event struct {
	ID      int
	Name    string    `orm:"size(100)"`
	From    string    `orm:"size(100)"`
	To      string    `orm:"size(100)"`
	Type    EventType `orm:"size(100)"`
}

func init(){
	orm.RegisterModel(new(Event))
}
