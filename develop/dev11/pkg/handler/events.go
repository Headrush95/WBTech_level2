package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	createEvent, err := decodeCreateInputToJson(r.Body)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	event := convertCreateInputToEventDomainModel(createEvent)

	h.service.Create(createEvent.UserId, event)

	SuccessPostDeleteResponse(w, fmt.Sprintf("[User %d] %d successfully created", createEvent.UserId, event.Id))
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	updateEvent, err := decodeUpdateInputToJson(r.Body)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Update(updateEvent)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessPostDeleteResponse(w, fmt.Sprintf("[User %d] %d successfully updated", updateEvent.UserId, updateEvent.EventId))
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		ErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	event, err := decodeDeleteInputToJson(r.Body)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Delete(event.UserId, event.EventId)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessPostDeleteResponse(w, fmt.Sprintf("[User %d] %d successfully deleted", event.UserId, event.EventId))
}

func (h *Handler) GetEventsPerDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, date, err := getUserIdAndDate(r.URL)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.service.GetEventsPerDay(userId, date)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessGetResponse(w, events)
}

func (h *Handler) GetEventsPerWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, date, err := getUserIdAndDate(r.URL)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.service.GetEventsPerWeek(userId, date)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessGetResponse(w, events)
}

func (h *Handler) GetEventsPerMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, date, err := getUserIdAndDate(r.URL)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.service.GetEventsPerMonth(userId, date)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	SuccessGetResponse(w, events)
}
