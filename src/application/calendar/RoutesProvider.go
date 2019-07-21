package calendar

import (
	"github.com/go-zoo/bone"
	"github.com/ugniusin/watchme-backend/src/application/calendar/controllers"
	"net/http"
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

	/*routes := []map[string]interface{}{
		{
			"method": "GET",
			"path": "/event",
			"controller": calendarController.ShowEvent,
		},
	}*/

	routesProvider.mux.Get("/calendar/:username", http.HandlerFunc(calendarController.Calendar))
	routesProvider.mux.Post("/event", http.HandlerFunc(calendarController.CreateEvent))
	routesProvider.mux.Get("/event/:id", http.HandlerFunc(calendarController.ShowEvent))
}
