package repository

import "rest/pkg/models"

type Event interface {
	CreateEvent(event models.Event) (int, error)
	UpdateEvent(event models.Event) (*models.Event, error)
	DeleteEvent(id int) error
	GetEvent(date1, date2 string) ([]models.Event, error)
}
type Repository struct {
	Event
}

func NewRepository(db *dBConnect) *Repository {
	return &Repository{Event: NewEventRepository(db)}
}
