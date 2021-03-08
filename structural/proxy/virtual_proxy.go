package proxy

type Subject interface {
	GetPath() string
	GetSize() int
	Display()
}

type Proxy struct {
	path string
	size int

	realSubject RealSubject
}

func (p *Proxy) GetPath() string {
	return p.path
}

func (p *Proxy) GetSize() int {
	return p.size
}

func (p *Proxy) Display() {
	p.realSubject.Display()
}

type RealSubject struct {

}

func (r *RealSubject) Display() {

}