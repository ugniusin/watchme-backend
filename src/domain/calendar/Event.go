package calendar

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Event struct {
	Id        int
	Name      string    `orm:"size(100)"`
	FromDate  time.Time `orm:"size(100)"`
	ToDate    time.Time `orm:"size(100)"`
	EventType EventType `orm:"size(100)"`
}

func init(){
	orm.RegisterModel(new(Event))
}

func NewEvent(
	name string,
	fromDate time.Time,
	toDate time.Time,
	eventType EventType,
) *Event {
	return &Event{
		Name:      name,
		FromDate:  fromDate,
		ToDate:    toDate,
		EventType: eventType,
	}
}
