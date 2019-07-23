package calendar

import (
	"github.com/go-zoo/bone"
	"github.com/ugniusin/watchme-backend/src/application/calendar/controllers"
)

type RoutesProvider struct {
	mux *bone.Mux
}

func NewRoutersProvider(mux *bone.Mux) *RoutesProvider {
	return &RoutesProvider{
		mux: mux,
	}
}

func (routesProvider *RoutesProvider) RegisterRoutes(calendarController *controllers.CalendarController) {

	routesProvider.mux.GetFunc("/calendar/:username", calendarController.Calendar)
	routesProvider.mux.PostFunc("/event", calendarController.CreateEvent)
	routesProvider.mux.GetFunc("/event/:id", calendarController.ShowEvent)
}
