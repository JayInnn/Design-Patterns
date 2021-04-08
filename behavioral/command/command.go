package command

type Command interface {
	Execute()
}

type AConcreteCommand struct {
	receiver Receiver
}

func (c *AConcreteCommand) Execute() {
	c.receiver.ConcreteActionA()
}

type BConcreteCommand struct {
	receiver Receiver
}

func (c *BConcreteCommand) Execute() {
	c.receiver.ConcreteActionB()
}
