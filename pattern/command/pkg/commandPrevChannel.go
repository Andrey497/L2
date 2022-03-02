package pkg

type PrevChannelCommand struct {
	Tv iTv
}

func (c *PrevChannelCommand) Execute() {
	c.Tv.prevChannel()
}
