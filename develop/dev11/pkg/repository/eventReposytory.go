package repository

import (
	"rest/pkg/models"
)

type eventRepository struct {
	db *dBConnect
}

func NewEventRepository(db *dBConnect) *eventRepository {
	return &eventRepository{db: db}
}

func (e *eventRepository) CreateEvent(event models.Event) (int, error) {
	queryString := `INSERT INTO event(name,date,description)  
					VALUES($1,$2,$3)`
	var id int
	err := e.db.QueryRow(queryString, event.Name, event.Date, event.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (e *eventRepository) UpdateEvent(event models.Event) (*models.Event, error) {
	queryString := `UPDATE event
    SET name=$1, date=$2, description=$3 
    WHERE id=$4
    RETURNING id, name, description, date`
	eventDB := new(models.EventDB)
	err := e.db.QueryRow(queryString, event.Name, event.Description, event.Date).Scan(eventDB.Id, eventDB.Name, eventDB.Date, eventDB.Description)
	if err != nil {
		return nil, err
	}
	return models.ConvertEventDbToEvent(eventDB), nil
}

func (e *eventRepository) DeleteEvent(id int) error {
	queryString := `Delete event
    				WHERE id=$1`
	_, err := e.db.Exec(queryString, id)
	if err != nil {
		return err
	}
	return nil
}
func (e *eventRepository) GetEvent(date1, date2 string) ([]models.Event, error) {
	queryString := `SELECT *
					FROM event
    				WHERE date BETWEEN $1 AND $2`
	rows, err := e.db.Query(queryString, date1, date2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []models.Event{}

	for rows.Next() {
		ev := models.EventDB{}
		err := rows.Scan(&ev.Id, &ev.Name, &ev.Date, &ev.Description)
		if err != nil {
			return nil, err
		}
		event := *models.ConvertEventDbToEvent(&ev)
		events = append(events, event)
	}
	return events, nil
}
