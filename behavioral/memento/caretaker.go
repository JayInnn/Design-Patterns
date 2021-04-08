package memento

// 负责人：对外只提供add和get操作，无法对备忘录进行操作和检查
type CareTaker struct {
	MementoList map[int]*Memento
}

func (ct *CareTaker) Add(index int, memento *Memento) {
	ct.MementoList[index] = memento
}

func (ct *CareTaker) Get(index int) *Memento {
	return ct.MementoList[index]
}