package bridge

type Color interface {
	PrintColor()
	DisplayColor()
}

// 红色逻辑
type Red struct {

}

func (r *Red) PrintColor() {

}

func (r *Red) DisplayColor() {

}

// 蓝色逻辑
type Blue struct {

}

func (b *Blue) PrintColor() {

}

func (b *Blue) DisplayColor() {

}