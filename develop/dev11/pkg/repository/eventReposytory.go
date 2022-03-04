package repository

type eventRepository struct {
	db *dBConnect
}

func NewEventRepository(db *dBConnect) *eventRepository {
	return &eventRepository{db: db}
}

func (e *eventRepository) CreateEvent(name, date, context string) (int, error) {
	queryString := `INSERT INTO event(name,date,description)  
					VALUES($1,$2,$3)`
	var id int
	err := e.db.QueryRow(queryString, name, date, context).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (e *eventRepository) UpdateEvent()     {}
func (e *eventRepository) DeleteEvent()     {}
func (e *eventRepository) GetEventByDay()   {}
func (e *eventRepository) GetEventByWeek()  {}
func (e *eventRepository) GetEventByMonth() {}
