package pkg

import "fmt"

type PowerOff struct {
	atm *atm
}

func (h *PowerOff) GiveMoney(int) error {
	return fmt.Errorf("it is not give money when atm is off")
}
func (h *PowerOff) AddMoney(int) error {
	return fmt.Errorf("it is not add money when atm is off")
}
func (h *PowerOff) PowerOn() error {
	h.atm.SetState(h.atm.powerOn)
	return nil
}
func (h *PowerOff) PowerOff() error {
	return nil
}
