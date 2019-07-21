package controllers

import (
	"encoding/json"
	"github.com/go-zoo/bone"
	"github.com/ugniusin/watchme-backend/src/domain/calendar"
	"github.com/ugniusin/watchme-backend/src/shared/application/errors"
	"net/http"
	"strconv"
	"time"
)

type CalendarController struct {
	eventRepository calendar.EventRepository
}

func NewCalendarController(
	eventRepository calendar.EventRepository,
) *CalendarController {
	return &CalendarController{
		eventRepository,
	}
}

func (controller *CalendarController) Calendar (w http.ResponseWriter, req *http.Request) {

	username := bone.GetValue(req, "username")

	js, err := json.Marshal(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (controller *CalendarController) ShowEvent (w http.ResponseWriter, req *http.Request) {

	eventIdParam := bone.GetValue(req, "id")

	eventId, err := strconv.Atoi(eventIdParam)
	if err != nil {
		err := errors.NewUnprocessableEntity("Invalid Event id")
		http.Error(w, err.Error(), err.GetCode())
		return
	}

	event, err := controller.eventRepository.FindOne(eventId); if err != nil {
		err := errors.NewNotFound(err.Error())
		http.Error(w, err.Error(), err.GetCode())
		return
	}

	js, err := json.Marshal(event); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (controller *CalendarController) CreateEvent (w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	type RequestBody struct {
		Name      string
		FromDate  string
		ToDate    string
		EventType string
	}

	var requestBody RequestBody

	err := decoder.Decode(&requestBody)
	if err != nil {
		panic(err)
	}

	layout := "2006-01-02 15:04:05"

	fromDate, err := time.Parse(layout, requestBody.FromDate)
	toDate, err := time.Parse(layout, requestBody.ToDate)

	event := calendar.NewEvent(
		requestBody.Name,
		fromDate,
		toDate,
		calendar.EventType(requestBody.EventType),
	)

	id := controller.eventRepository.Save(*event)

	response := map[string]interface{}{
		"Id": id,
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
