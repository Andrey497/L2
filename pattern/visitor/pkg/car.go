package pkg

type car struct {
	model string
	sped  int
}

func (c *car) Accept(v visitor) {
	v.visitorCar(c)
}

func (c *car) GetType() string {
	return "car"
}
func NewCar() *car {
	return &car{}
}
