package service

import (
	"dev11/models"
	"dev11/pkg/repository"
	"time"
)

type Service struct {
	PostEvents
	GetEvents
}

func NewService(repo *repository.Repository) *Service {
	return &Service{PostEvents: repo.CacheRepository, GetEvents: repo.CacheRepository}
}

type PostEvents interface {
	Create(userId int, event models.Event)
	Update(event models.UpdateEventModel) error
	Delete(userId int, eventId uint32) error
}

type GetEvents interface {
	GetEventsPerDay(userId int, day time.Time) ([]models.Event, error)
	GetEventsPerWeek(userId int, day time.Time) ([]models.Event, error)
	GetEventsPerMonth(userId int, day time.Time) ([]models.Event, error)
}
