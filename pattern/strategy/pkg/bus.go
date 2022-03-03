package pkg

import "fmt"

type bus struct {
}

func NewBus() *bus {
	return &bus{}
}
func (b *bus) Move(u *user) {
	fmt.Println(fmt.Sprintf("User %v moving from %v to %v by bus", u.userName, u.GetAddressFrom(), u.GetAddressTo()))
}
