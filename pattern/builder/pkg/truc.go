package pkg

type trucCarBuilder struct {
	typeCar  string
	weight   int
	maxSpeed int
}

func newTrucBuilder() *trucCarBuilder {
	return &trucCarBuilder{}
}
func (t *trucCarBuilder) setWeight() {
	t.weight = 123
}
func (t *trucCarBuilder) setType() {
	t.typeCar = "passengerCar"
}
func (t *trucCarBuilder) seMaxSpeed() {
	t.maxSpeed = 90
}

func (t *trucCarBuilder) getCar() *car {
	return &car{
		typeCar:  t.typeCar,
		weight:   t.weight,
		maxSpeed: t.maxSpeed,
	}
}
