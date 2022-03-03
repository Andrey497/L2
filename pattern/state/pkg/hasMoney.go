package pkg

import "fmt"

type HasMoney struct {
	atm *atm
}

func (h *HasMoney) GiveMoney(money int) error {
	if h.atm.moneyInATM < money {
		return fmt.Errorf("there is not enough money in the ATM")
	}
	h.atm.moneyInATM -= money
	if h.atm.moneyInATM == 0 {
		h.atm.SetState(h.atm.noMoney)
	}
	return nil

}
func (h *HasMoney) AddMoney(money int) error {
	h.AddMoney(money)
	h.atm.SetState(h.atm.hasMoney)
	return nil
}
func (h *HasMoney) PowerOn() error {
	h.atm.SetState(h.atm.powerOn)
	return nil
}
func (h *HasMoney) PowerOff() error {
	h.atm.SetState(h.atm.powerOff)
	return nil
}
