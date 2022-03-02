package pkg

type ship struct {
	model string
	sped  int
}

func (s *ship) Accept(v visitor) {
	v.visitorShip(s)
}

func (s *ship) GetType() string {
	return "ship"
}
func NewShip() *ship {
	return &ship{}
}
