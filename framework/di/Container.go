package di

import (
	"github.com/ugniusin/watchme/framework"
	"github.com/ugniusin/watchme/framework/config"
	"github.com/ugniusin/watchme/framework/database"
	"github.com/ugniusin/watchme/src/application/calendar/controllers"
	"github.com/ugniusin/watchme/src/domain/calendar"
	calendar2 "github.com/ugniusin/watchme/src/infrastructure/calendar"
	"go.uber.org/dig"
)

type Container struct {
}


func BuildContainer() *dig.Container {
	container := dig.New()

	//calendar.BuildContainer(container)

	container.Provide(config.NewConfiguration)
	container.Provide(database.NewDatabase)

	container.Provide(controllers.NewCalendarController)
	container.Provide(calendar.NewEventCreator)
	container.Provide(calendar2.NewEventRepository, dig.As(new(calendar.EventRepository)))

	container.Provide(framework.NewFrontController)
	container.Provide(framework.NewServer)

	return container
}
