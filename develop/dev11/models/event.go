package models

import (
	"encoding/json"
	"strings"
	"time"
)

type JSONTime time.Time

func (jt *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*jt = JSONTime(t)
	return nil
}

func (jt *JSONTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*jt))
}

type Event struct {
	Id          uint32   `json:"id"`
	Date        JSONTime `json:"date"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
}

type CreateEventModel struct {
	UserId      int      `json:"user_id" validate:"required"`
	Date        JSONTime `json:"date" validate:"required"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description"`
}

type DeleteEventModel struct {
	UserId  int    `json:"user_id" validate:"required"`
	EventId uint32 `json:"event_id" validate:"required"`
}

type UpdateEventModel struct {
	UserId      int      `json:"user_id" validate:"required"`
	EventId     uint32   `json:"event_id" validate:"required"`
	Date        JSONTime `json:"date"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
}
