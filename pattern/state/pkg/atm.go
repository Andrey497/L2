package pkg

type atm struct {
	hasMoney state
	noMoney  state
	powerOn  state
	powerOff state

	currentState state
	moneyInATM   int
}

func NewAtm(money int) *atm {
	a := &atm{moneyInATM: money}
	hasMoney := &HasMoney{a}
	noMoney := &NoMoney{a}
	powerOn := &PowerOn{a}
	powerOff := &PowerOff{a}

	a.hasMoney = hasMoney
	a.noMoney = noMoney
	a.powerOn = powerOn
	a.powerOff = powerOff

	a.SetState(powerOff)
	return a
}

func (a *atm) SetState(s state) {
	a.currentState = s
}

func (a *atm) GiveMoney(money int) error {
	return a.currentState.GiveMoney(money)
}
func (a *atm) AddMoney(money int) error {
	return a.currentState.AddMoney(money)
}
func (a *atm) PowerOn() error {
	return a.currentState.PowerOn()
}
func (a *atm) PowerOff() error {
	return a.currentState.PowerOff()
}
