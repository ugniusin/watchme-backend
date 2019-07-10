package controllers

import (
	"encoding/json"
	"github.com/go-zoo/bone"
	"github.com/ugniusin/watchme/framework/config"
	"github.com/ugniusin/watchme/src/domain/calendar"
	"github.com/ugniusin/watchme/src/domain/calendar/dto"
	"net/http"
)

type CalendarController struct {
	ClientId string
	ClientSecret string
	RedirectUri string
	EventCreator *calendar.EventCreator
}

func NewCalendarController(
	config *config.Configuration,
	eventCreator *calendar.EventCreator,
) *CalendarController {
	return &CalendarController{
		config.Instagram["ClientId"],
		config.Instagram["ClientSecret"],
		config.Instagram["RedirectUri"],
		eventCreator,
	}
}

func (controller *CalendarController) Calendar (w http.ResponseWriter, req *http.Request)  {

	username := bone.GetValue(req, "username")

	js, err := json.Marshal(dto.User{
		Username: username,
		FullName: "Random Full Name",
		ProfilePicture: "Random Picture",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (controller *CalendarController) InsertUser (w http.ResponseWriter, req *http.Request)  {

	controller.EventCreator.Create()

	js, err := json.Marshal(dto.User{
		Username: "Random Username",
		FullName: "Random Full Name",
		ProfilePicture: "Random Picture",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
