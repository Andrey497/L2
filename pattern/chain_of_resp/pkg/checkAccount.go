package pkg

import "errors"

type CheckAccount struct {
	Next     iOperation
	Login    string
	Password string
	Err      error
}

func (c *CheckAccount) Execute(a Account) {
	if c.Login != a.Login || c.Password != a.Password {
		c.Err = errors.New("Ошибка: Неверный логин или пароль")
		c.Next = nil

	} else {
		if c.Next != nil {
			c.Next.Execute(a)
		}
	}
}
func (c *CheckAccount) SetNext(nextOperation iOperation) {
	c.Next = nextOperation
}
