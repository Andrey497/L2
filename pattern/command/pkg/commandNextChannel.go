package pkg

type NextChannelCommand struct {
	Tv iTv
}

func (c *NextChannelCommand) Execute() {
	c.Tv.nextChannel()
}
