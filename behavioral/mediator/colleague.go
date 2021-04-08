package mediator

import "fmt"

type Colleague interface {
	SetMediator(mediator Mediator)
	SendMessage(message string)
	GetMessage(message string)
}

// 具体同事A
type ConcreteColleagueA struct {
	mediator Mediator
}

func (a ConcreteColleagueA) SetMediator(mediator Mediator) {
	a.mediator = mediator
}

func (a ConcreteColleagueA) SendMessage(message string) {
	a.mediator.ForwardMessage(message, a)
}

func (a ConcreteColleagueA) GetMessage(message string) {
	fmt.Println("get message: " + message)
}

// 具体同事B
type ConcreteColleagueB struct {
	mediator Mediator
}

func (b ConcreteColleagueB) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

func (b ConcreteColleagueB) SendMessage(message string) {
	b.mediator.ForwardMessage(message, b)
}

func (b ConcreteColleagueB) GetMessage(message string) {
	fmt.Println("get message: " + message)
}
