package framework

import (
	"github.com/go-zoo/bone"
	"github.com/ugniusin/watchme/framework/config"
	"net/http"
)

type Server struct {
	configuration *config.Configuration
	frontController *FrontController
}

func NewServer(
	configuration *config.Configuration,
	frontController *FrontController,
) *Server {

	return &Server{
		configuration: configuration,
		frontController: frontController,
	}
}

func (s *Server) Run() {

	//http.HandleFunc("/calendar", s.frontController.calendarController.Calendar)
	//http.HandleFunc("/insert", s.frontController.calendarController.InsertUser)

	mux := bone.New()

	mux.Get("/calendar/:username", http.HandlerFunc(s.frontController.calendarController.Calendar))
	mux.Get("/insert", http.HandlerFunc(s.frontController.calendarController.InsertUser))

	http.ListenAndServe(":8090", mux)



	//http.ListenAndServe(":8090", nil)
}