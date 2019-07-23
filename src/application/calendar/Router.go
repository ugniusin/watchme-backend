package calendar

import (
	"github.com/go-zoo/bone"
	"github.com/ugniusin/watchme-backend/src/application/calendar/controllers"
)

type Router struct {
	mux *bone.Mux
}

func NewRouter(mux *bone.Mux) *Router {
	return &Router{
		mux: mux,
	}
}

func (router *Router) RegisterRoutes(calendarController *controllers.CalendarController) {

	router.mux.GetFunc("/calendar/:username", calendarController.Calendar)
	router.mux.PostFunc("/event", calendarController.CreateEvent)
	router.mux.GetFunc("/event/:id", calendarController.ShowEvent)
}
