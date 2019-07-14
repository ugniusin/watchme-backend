package framework

import (
	"github.com/go-zoo/bone"
	"github.com/ugniusin/watchme-backend/framework/config"
	"github.com/ugniusin/watchme-backend/framework/database"
	"net/http"
)

type Server struct {
	configuration *config.Configuration
	frontController *FrontController
	database *database.Database
}

func NewServer(
	configuration *config.Configuration,
	frontController *FrontController,
	database *database.Database,
) *Server {

	return &Server{
		configuration: configuration,
		frontController: frontController,
		database: database,
	}
}

func (s *Server) Run() {

	s.database.BootstrapDatabase()

	mux := bone.New()

	mux.Get("/calendar/:username", http.HandlerFunc(s.frontController.calendarController.Calendar))
	mux.Post("/event", http.HandlerFunc(s.frontController.calendarController.CreateEvent))

	http.ListenAndServe(":8090", mux)
}