package di

import (
	"github.com/ugniusin/watchme-backend/framework"
	config2 "github.com/ugniusin/watchme-backend/framework/config"
	database2 "github.com/ugniusin/watchme-backend/framework/database"
	"github.com/ugniusin/watchme-backend/src/application/calendar"
	"go.uber.org/dig"
)

type Container struct {
}


func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(config2.NewConfiguration)
	container.Provide(database2.NewDatabase)
	bootstrapContexts(container)
	container.Provide(framework.NewFrontController)
	container.Provide(framework.NewRouter)
	container.Provide(framework.NewServer)

	return container
}

func bootstrapContexts(container *dig.Container) {
	calendar.Bootstrap(container)
}