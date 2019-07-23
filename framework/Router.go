package framework

import (
	"github.com/go-zoo/bone"
	"github.com/ugniusin/watchme-backend/src/application/calendar"
)

type Router struct {
	frontController *FrontController
}

func NewRouter(
	frontController *FrontController,
) *Router {

	return &Router{
		frontController: frontController,
	}
}

func (router *Router) RegisterRoutes(mux *bone.Mux) {

	calendarRouter := calendar.NewRouter(mux)
	calendarRouter.RegisterRoutes(router.frontController.calendarController)
}