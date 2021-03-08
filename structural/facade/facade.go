package facade

type Facade struct {
	graph Graph
	color Color
}

func NewFacade(graph Graph, color Color) *Facade {
	return &Facade{
		graph: graph,
		color: color,
	}
}

func (f *Facade) DrawGraph() {
	f.graph.Draw()
}

func (f *Facade) PrintColor() {
	f.color.Print()
}