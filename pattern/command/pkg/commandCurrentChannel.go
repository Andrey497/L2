package pkg

type CurrentChannelCommand struct {
	Tv iTv
}

func (c *CurrentChannelCommand) Execute() {
	c.Tv.currentChannel()
}
