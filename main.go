package main

import (
	"github.com/ugniusin/watchme/framework"
	"github.com/ugniusin/watchme/framework/database"
	"github.com/ugniusin/watchme/framework/di"
)

func main() {


	database.NewDatabase()

	container := di.BuildContainer()

	err := container.Invoke(func(server *framework.Server) {

		server.Run()
	})

	if err != nil {
		panic(err)
	}

}
