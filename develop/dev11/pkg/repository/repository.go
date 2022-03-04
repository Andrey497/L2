package repository

type Event interface {
	CreateEvent()
	UpdateEvent()
	DeleteEvent()
	GetEventByDay()
	GetEventByWeek()
	GetEventByMonth()
}
type repository struct {
	Event Event
}

func NewRepository(db *dBConnect) *repository {
	return &repository{Event: NewEventRepository(db)}
}
