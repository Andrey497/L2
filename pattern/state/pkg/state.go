package pkg

type state interface {
	GiveMoney(int) error
	AddMoney(int) error
	PowerOn() error
	PowerOff() error
}
