package mediator

type Mediator interface {
	ForwardMessage(message string, colleague Colleague)
}

type ConcreteMediator struct {
	colleagueA ConcreteColleagueA
	colleagueB ConcreteColleagueB
}

func (c *ConcreteMediator) ForwardMessage(message string, colleague Colleague) {
	switch colleague.(type) {
	case ConcreteColleagueA:
		c.colleagueB.GetMessage(message)
	case ConcreteColleagueB:
		c.colleagueA.GetMessage(message)
	}
}