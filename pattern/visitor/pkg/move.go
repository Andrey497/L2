package pkg

import "fmt"

type move struct {
}

func NewMove() move {
	return move{}
}

func (m *move) visitorCar(c *car) {
	fmt.Println("car move")
}
func (m *move) visitorPlane(p *plane) {
	fmt.Println("Plane move")
}
func (m *move) visitorShip(s *ship) {
	fmt.Println("Ship move")
}
