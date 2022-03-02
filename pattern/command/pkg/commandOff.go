package pkg

type OffCommand struct {
	Tv iTv
}

func (c *OffCommand) Execute() {
	c.Tv.tvOff()
}
