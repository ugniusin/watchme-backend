package calendar

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	calendar2 "github.com/ugniusin/watchme-backend/src/domain/calendar"
	"log"
)

type EventRepository struct {
}

func NewEventRepository() *EventRepository {
	return &EventRepository{}
}

func (eventRepository *EventRepository) Save(event calendar2.Event) {

	o := orm.NewOrm()
	o.Using("default")

	var res int64
	var err error

	// Insert
	if res, err = o.Insert(&event); err != nil {
		log.Println(err)
	}
	log.Printf("inserted: %d row", res)
}