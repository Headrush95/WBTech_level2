package handler

import (
	"dev11/pkg/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

/*
POST /create_event
● POST /update_event
● POST /delete_event
● GET /events_for_day
● GET /events_for_week
● GET /events_for_month
*/

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)
	mux.HandleFunc("/events_for_day", h.GetEventsPerDay)
	mux.HandleFunc("/events_for_week", h.GetEventsPerWeek)
	mux.HandleFunc("/events_for_month", h.GetEventsPerMonth)

	return logger(mux)
}
