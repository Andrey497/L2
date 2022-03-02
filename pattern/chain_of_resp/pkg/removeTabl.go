package pkg

import (
	"fmt"
)

type RemoveTable struct {
	Next iOperation
}

func (c *RemoveTable) Execute(a Account) {
	fmt.Println("remove table")
	if c.Next != nil {
		c.Next.Execute(a)
	}
}
func (c *RemoveTable) SetNext(nextOperation iOperation) {
	c.Next = nextOperation
}
