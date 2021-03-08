package facade

type Color interface {
	Print()
}

type Red struct {

}

func (r *Red) Print() {

}

type Blue struct {

}

func (b *Blue) Print() {

}

type Yellow struct {

}

func (y *Yellow) Print() {

}