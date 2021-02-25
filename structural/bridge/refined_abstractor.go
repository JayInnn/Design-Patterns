package bridge

// refined graph
type RefinedGraph interface {
	Graph
	RefinedOperation()
}

type RefinedCircle struct {
	color Color
}

func (re *RefinedCircle) Operation() {
	re.color.PrintColor()
	re.color.DisplayColor()
}

func (re *RefinedCircle) RefinedOperation() {

}