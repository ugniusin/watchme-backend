package calendar

import (
	"github.com/ugniusin/watchme-backend/src/application/calendar/controllers"
	"github.com/ugniusin/watchme-backend/src/domain/calendar"
	calendar2 "github.com/ugniusin/watchme-backend/src/infrastructure/calendar"
	"go.uber.org/dig"
)

func Bootstrap(container *dig.Container) {

	container.Provide(controllers.NewCalendarController)
	container.Provide(calendar2.NewEventRepository, dig.As(new(calendar.EventRepository)))
}
