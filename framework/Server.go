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
	router *Router
	database *database.Database
}

func NewServer(
	configuration *config.Configuration,
	frontController *FrontController,
	router *Router,
	database *database.Database,
) *Server {

	return &Server{
		configuration: configuration,
		frontController: frontController,
		router: router,
		database: database,
	}
}

func (s *Server) Run() {

	s.database.BootstrapDatabase()

	mux := bone.New()

	//someSlice := [3]string{"Apple", "Mango", "Banana"}

	/*for index, element := range someSlice {
		// index is the index where we are
		// element is the element from someSlice for where we are
	}*/

	s.router.RegisterRoutes(mux)


	http.ListenAndServe(":8090", mux)
}