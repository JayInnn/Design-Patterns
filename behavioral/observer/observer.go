package observer

type Observer interface {
	Update()
}

type ConcreteObserver struct {
	subject Subject
}

func NewConcreteObserver(subject Subject) *ConcreteObserver {
	co := &ConcreteObserver{
		subject: subject,
	}
	co.subject.Attach(co)
	return co
}

func (co *ConcreteObserver) Update() {
	//TODO: 观察subject更新后的操作
}