package bridge

// 提供client接口
type Graph interface {
	Operation()
}

// 圆形
type Circle struct {
	color Color
}

func (circle *Circle) Operation() {
	circle.color.PrintColor()
	circle.color.DisplayColor()
}

// 三角形
type Triangle struct {
	color Color
}

func (triangle *Triangle) Operation() {
	triangle.color.PrintColor()
	triangle.color.DisplayColor()
}
