package observer

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
}

type ConcreteSubject struct {
	obList []Observer
	state int
}

// 状态变更，通知observer
func (cs *ConcreteSubject) SetState(st int) {
	if cs.state != st {
		cs.state = st
		cs.Notify()
	}
}

// 注册observer
func (cs *ConcreteSubject) Attach(observer Observer) {
	cs.obList = append(cs.obList, observer)
}

// 删除observer
func (cs *ConcreteSubject) Detach(observer Observer) {

}

func (cs *ConcreteSubject) Notify() {
	for _, observer := range cs.obList{
		observer.Update()
	}
}