package command

type Command interface {
	Execute()
}

type OnCommand struct {
	D Device
}

func (c *OnCommand) Execute() {
	c.D.On()
}

type OffCommand struct {
	D Device
}

func (c *OffCommand) Execute() {
	c.D.Off()
}
