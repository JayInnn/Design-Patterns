package bridge

// factory-method 模式创建
type GraphFactory interface {
	createGraph() Graph
}

// 类circle实例
type CircleFactory struct {

}

func (c *CircleFactory) createGraph() Graph {
	return new(Circle)
}

// 类Triangle实例
type TriangleFactory struct {

}

func (t *TriangleFactory) createGraph() Graph {
	return new(Triangle)
}