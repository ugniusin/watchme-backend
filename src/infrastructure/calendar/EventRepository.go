package calendar

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	domainCalendar "github.com/ugniusin/watchme-backend/src/domain/calendar"
	"github.com/ugniusin/watchme-backend/src/shared/infrastructure/errors"
	"log"
)

type EventRepository struct {
}

func NewEventRepository() *EventRepository {
	return &EventRepository{}
}

func (eventRepository *EventRepository) FindOne (id int) (domainCalendar.Event, error) {

	o := orm.NewOrm()

	// read one
	event := domainCalendar.Event{Id: id}

	var err error

	if err = o.Read(&event); err != nil {
		err := errors.NewError("Event not found")

		return domainCalendar.Event{}, err
	}

	return event, nil
}

func (eventRepository *EventRepository) Save (event domainCalendar.Event) int {

	o := orm.NewOrm()

	var res int64
	var err error

	// Insert
	if res, err = o.Insert(&event); err != nil {
		log.Println(err)
	}
	log.Printf("inserted: %d row", res)

	return event.Id
}
