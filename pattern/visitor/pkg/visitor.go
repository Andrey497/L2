package pkg

type visitor interface {
	visitorCar(c *car)
	visitorPlane(p *plane)
	visitorShip(s *ship)
}
