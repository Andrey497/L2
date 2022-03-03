package pkg

import "fmt"

type tax struct {
}

func NewTax() *tax {
	return &tax{}
}
func (t *tax) Move(u *user) {
	fmt.Println(fmt.Sprintf("User %v moving from %v to %v by tax", u.userName, u.GetAddressFrom(), u.GetAddressTo()))
}
