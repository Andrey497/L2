package pkg

type CheckIsAdministrator struct {
	Next            iOperation
	IsAdministrator bool
}

func (c *CheckIsAdministrator) Execute(a Account) {
	if a.IsAdministrator != true {
		c.IsAdministrator = false
		c.Next = nil
	} else {
		c.IsAdministrator = true
		if c.Next != nil {
			c.Next.Execute(a)
		}
	}
}
func (c *CheckIsAdministrator) SetNext(nextOperation iOperation) {
	c.Next = nextOperation
}
