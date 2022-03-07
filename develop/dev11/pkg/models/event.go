package models

import "time"

type Event struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type EventDB struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

func ConvertEventDbToEvent(e *EventDB) *Event {
	return &Event{
		Id:          e.Id,
		Name:        e.Name,
		Date:        e.Date.Format("2006-01-02"),
		Description: e.Description,
	}
}

func (e *Event) Valid() bool {
	if e.Name != "" {
		_, err := time.Parse("2006-01-02", e.Date)
		if err == nil {
			return true
		}
	}
	return false
}
