package pkg

import "fmt"

type walking struct {
}

func NewWalking() *walking {
	return &walking{}
}
func (w *walking) Move(u *user) {
	fmt.Println(fmt.Sprintf("User %v moving from %v to %v by walking", u.userName, u.GetAddressFrom(), u.GetAddressTo()))
}
