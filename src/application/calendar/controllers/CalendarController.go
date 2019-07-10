package controllers

import (
	"encoding/json"
	"github.com/go-zoo/bone"
	"github.com/ugniusin/watchme/src/domain/calendar"
	"net/http"
)

type CalendarController struct {
	EventCreator *calendar.EventCreator
}

func NewCalendarController(
	eventCreator *calendar.EventCreator,
) *CalendarController {
	return &CalendarController{
		eventCreator,
	}
}

func (controller *CalendarController) Calendar (w http.ResponseWriter, req *http.Request)  {

	username := bone.GetValue(req, "username")

	js, err := json.Marshal(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (controller *CalendarController) InsertUser (w http.ResponseWriter, req *http.Request)  {

	controller.EventCreator.Create()

	js, err := json.Marshal(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
