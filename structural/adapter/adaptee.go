package adapter

type Adaptee struct {

}

// 具体实现
func (a *Adaptee) SpecificRequest() string {
	return "this is specific request!"
}
