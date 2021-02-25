package adapter

type Adapter struct {
	adaptee Adaptee
}

func (adapter *Adapter) Request() {
	adapter.adaptee.SpecificRequest()
}
