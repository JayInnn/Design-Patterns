package strategy

type StrategyContext interface {
	ContextInterface()
}

type ConcreteStrategyContext struct {
	strategy Strategy
}

func (c *ConcreteStrategyContext) ContextInterface() {

}


