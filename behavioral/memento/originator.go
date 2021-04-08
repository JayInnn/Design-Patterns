package memento

type Originator struct {
	state string
}

func (o *Originator) SetMemento(memento Memento) {
	o.state = memento.state
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{
		state: o.state,
	}
}