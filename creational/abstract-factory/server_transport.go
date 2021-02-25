package abstract_factory

type ServerTransport interface {
	ListenAndServe()
}

type TCPTransport struct {

}

func (t *TCPTransport) ListenAndServe() {

}

type UDPTransport struct {

}

func (u *UDPTransport) ListenAndServe() {

}


