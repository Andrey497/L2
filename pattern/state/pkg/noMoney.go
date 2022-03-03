package pkg

import "fmt"

type NoMoney struct {
	atm *atm
}

func (h *NoMoney) GiveMoney(int) error {
	return fmt.Errorf("insufficient funds")
}
func (h *NoMoney) AddMoney(money int) error {
	h.atm.moneyInATM += money
	h.atm.SetState(h.atm.hasMoney)
	return nil
}
func (h *NoMoney) PowerOn() error {
	h.atm.SetState(h.atm.powerOn)
	return nil
}
func (h *NoMoney) PowerOff() error {
	h.atm.SetState(h.atm.powerOff)
	return nil
}
