package memento

type Memento struct {
	state string
}

func (m *Memento) SetState(st string) {
	m.state = st
}

func (m *Memento) GetState() string {
	return m.state
}
