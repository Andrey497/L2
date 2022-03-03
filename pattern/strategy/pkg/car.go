package pkg

import "fmt"

type car struct {
}

func NewCar() *car {
	return &car{}
}
func (c *car) Move(u *user) {
	fmt.Println(fmt.Sprintf("User %v moving from %v to %v by car", u.userName, u.GetAddressFrom(), u.GetAddressTo()))
}
