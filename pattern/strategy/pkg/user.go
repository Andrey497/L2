package pkg

type user struct {
	userName    string
	addressTo   string
	addressFrom string
	typeMoving  iMoving
}

func NewUser(userName string) *user {
	return &user{userName: userName}
}

func (u *user) SetRout(addressFrom string, addressTo string) {
	u.addressTo = addressTo
	u.addressFrom = addressFrom
}
func (u *user) GetAddressTo() string {
	return u.addressTo
}
func (u *user) GetAddressFrom() string {
	return u.addressFrom
}
func (u *user) SetTypeMoving(typeMoving iMoving) {
	u.typeMoving = typeMoving
}
func (u *user) Moving() {
	u.typeMoving.Move(u)
}
