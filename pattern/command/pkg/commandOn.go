package pkg

type OnCommand struct {
	Tv iTv
}

func (c *OnCommand) Execute() {
	c.Tv.tvOn()
}
