package strategy

type Strategy interface {
	AlgorithmInterface()
}

// 具体算法A
type ConcreteStrategyA struct {

}

func (cs *ConcreteStrategyA) AlgorithmInterface() {

}

// 具体算法B
type ConcreteStrategyB struct {

}

func (cs *ConcreteStrategyB) AlgorithmInterface() {

}