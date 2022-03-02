package pkg

type account struct {
	id          int
	nameAccount string
	phone       string
}

func createAccount(id int, nameAccount string, phone string) *account {
	return &account{
		id:          id,
		nameAccount: nameAccount,
		phone:       phone,
	}
}

func (a *account) checkAccountByPhone(phone string) bool {
	if a.phone == phone {
		return true
	}
	return false
}
