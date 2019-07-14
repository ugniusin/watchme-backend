package framework

import (
	calendarControllers "github.com/ugniusin/watchme-backend/src/application/calendar/controllers"
)

type FrontController struct {
	calendarController *calendarControllers.CalendarController
}

func NewFrontController(calendarController *calendarControllers.CalendarController) *FrontController {

	return &FrontController{
		calendarController: calendarController,
	}
}
