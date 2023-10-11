package repository

import (
	"dev11/models"
	"errors"
	"slices"
	"sync"
	"time"
)

var (
	defaultEventSliceLength = 10
	notFoundEvent           = errors.New("event was not found")
	notFoundUser            = errors.New("user was not found")
)

type Cache struct {
	mu   sync.RWMutex
	data map[int][]models.Event
}

func NewCache(size int) *Cache {
	return &Cache{mu: sync.RWMutex{}, data: make(map[int][]models.Event, size)}
}

func searchEventsInInterval(events []models.Event, day time.Time, interval string) []models.Event {
	result := make([]models.Event, 0, defaultEventSliceLength)
	var timeInFuture time.Time
	switch interval {
	case "day":
		timeInFuture = day.Add(time.Second)
	case "week":
		timeInFuture = day.Add(7 * 24 * time.Hour)
	case "month":
		timeInFuture = day.Add(30 * 24 * time.Hour)
	default:
		timeInFuture = day.Add(time.Second)
	}

	for _, event := range events {
		if time.Time(event.Date).Before(timeInFuture) && time.Time(event.Date).After(day) || time.Time(event.Date).Equal(day) {
			result = append(result, event)
		}
	}
	return result
}

func (c *Cache) Create(userId int, event models.Event) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[userId]; !ok {
		c.data[userId] = make([]models.Event, 0, defaultEventSliceLength)
	}

	c.data[userId] = append(c.data[userId], event)
}

func (c *Cache) Update(UpdatedEvent models.UpdateEventModel) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[UpdatedEvent.UserId]; !ok {
		return notFoundUser
	}

	foundIdx := -1
	for eventIdx, event := range c.data[UpdatedEvent.UserId] {
		if event.Id == UpdatedEvent.EventId {
			foundIdx = eventIdx
			break
		}
	}

	if foundIdx == -1 {
		return notFoundEvent
	}

	if !time.Time(UpdatedEvent.Date).Equal(time.Time{}) {
		c.data[UpdatedEvent.UserId][foundIdx].Date = UpdatedEvent.Date
	}

	if UpdatedEvent.Title != "" {
		c.data[UpdatedEvent.UserId][foundIdx].Title = UpdatedEvent.Title
	}

	if UpdatedEvent.Description != "" {
		c.data[UpdatedEvent.UserId][foundIdx].Description = UpdatedEvent.Description
	}

	return nil
}

func (c *Cache) Delete(userId int, eventId uint32) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[userId]; !ok {
		return notFoundUser
	}

	found := false
	for id, event := range c.data[userId] {
		if event.Id == eventId {
			found = true
			slices.Delete(c.data[userId], id, id+1)
			break
		}
	}

	if !found {
		return notFoundEvent
	}
	return nil
}

func (c *Cache) GetEventsPerDay(userId int, day time.Time) ([]models.Event, error) {
	if _, ok := c.data[userId]; !ok {
		return nil, notFoundUser
	}

	return searchEventsInInterval(c.data[userId], day, "day"), nil
}

func (c *Cache) GetEventsPerWeek(userId int, day time.Time) ([]models.Event, error) {
	if _, ok := c.data[userId]; !ok {
		return nil, notFoundUser
	}

	return searchEventsInInterval(c.data[userId], day, "week"), nil
}

func (c *Cache) GetEventsPerMonth(userId int, day time.Time) ([]models.Event, error) {
	if _, ok := c.data[userId]; !ok {
		return nil, notFoundUser
	}

	return searchEventsInInterval(c.data[userId], day, "month"), nil
}
