package pkg

type passengerCarBuilder struct {
	typeCar  string
	weight   int
	maxSpeed int
}

func newPassengerCarBuilder() *passengerCarBuilder {
	return &passengerCarBuilder{}
}
func (p *passengerCarBuilder) setWeight() {
	p.weight = 123
}
func (p *passengerCarBuilder) setType() {
	p.typeCar = "passengerCar"
}
func (p *passengerCarBuilder) seMaxSpeed() {
	p.maxSpeed = 90
}

func (p *passengerCarBuilder) getCar() *car {
	return &car{
		typeCar:  p.typeCar,
		weight:   p.weight,
		maxSpeed: p.maxSpeed,
	}
}
