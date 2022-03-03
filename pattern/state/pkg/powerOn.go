package pkg

import "fmt"

type PowerOn struct {
	atm *atm
}

func (h *PowerOn) GiveMoney(money int) error {
	if h.atm.moneyInATM < money {
		return fmt.Errorf("there is not enough money in the ATM")
	}
	h.atm.moneyInATM -= money
	if h.atm.moneyInATM == 0 {
		h.atm.SetState(h.atm.noMoney)
	}
	return nil
}
func (h *PowerOn) AddMoney(money int) error {
	h.AddMoney(money)
	h.atm.SetState(h.atm.hasMoney)
	return nil
}
func (h *PowerOn) PowerOn() error {
	return nil
}
func (h *PowerOn) PowerOff() error {
	h.atm.SetState(h.atm.powerOn)
	return nil
}
