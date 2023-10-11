package handler

import (
	"dev11/models"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Printf("[%s] %s\n", r.Method, r.URL)

	})
}

func decodeCreateInputToJson(input io.ReadCloser) (models.CreateEventModel, error) {
	var event models.CreateEventModel
	err := json.NewDecoder(input).Decode(&event)

	return event, err
}

func decodeUpdateInputToJson(input io.ReadCloser) (models.UpdateEventModel, error) {
	var event models.UpdateEventModel
	err := json.NewDecoder(input).Decode(&event)

	return event, err
}

func decodeDeleteInputToJson(input io.ReadCloser) (models.DeleteEventModel, error) {
	var event models.DeleteEventModel
	err := json.NewDecoder(input).Decode(&event)

	return event, err
}

func getUserIdAndDate(input *url.URL) (int, time.Time, error) {
	userId, err := strconv.Atoi(input.Query().Get("user_id"))
	if err != nil {
		return 0, time.Time{}, err
	}

	date, err := time.Parse(time.DateOnly, input.Query().Get("date"))
	return userId, date, err
}

func convertCreateInputToEventDomainModel(input models.CreateEventModel) models.Event {
	var event models.Event

	idRef, _ := uuid.NewUUID()

	event.Id = idRef.ID()
	event.Date = input.Date
	event.Title = input.Title
	event.Description = input.Description

	return event
}
