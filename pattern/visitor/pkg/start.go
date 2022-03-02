package pkg

import "fmt"

type start struct {
}

func (s *start) visitorForCar(c *car) {
	fmt.Println("car start")
}
func (s *start) visitorPlane(p *plane) {
	fmt.Println("Plane start")
}
func (s *start) visitorShip(sh *ship) {
	fmt.Println("Ship start")
}
