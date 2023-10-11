package repository

import (
	"dev11/models"
	"time"
)

type Repository struct {
	CacheRepository
}

func NewRepository(cache *Cache) *Repository {
	return &Repository{CacheRepository: cache}
}

type CacheRepository interface {
	Create(userId int, event models.Event)
	Update(event models.UpdateEventModel) error
	Delete(userId int, eventId uint32) error
	GetEventsPerDay(userId int, day time.Time) ([]models.Event, error)
	GetEventsPerWeek(userId int, day time.Time) ([]models.Event, error)
	GetEventsPerMonth(userId int, day time.Time) ([]models.Event, error)
}
