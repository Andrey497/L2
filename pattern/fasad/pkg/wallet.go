package pkg

type wallet struct {
	id      int
	balance int
}

func newWallet(id int) *wallet {
	return &wallet{
		id:      id,
		balance: 0,
	}
}
func (w *wallet) AddMoney(money int) {
	w.balance += money
}
func (w *wallet) CheckBalance() int {
	return w.balance
}
