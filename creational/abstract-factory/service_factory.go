package abstract_factory

type AbstractServiceFactory interface {
	CreateServerTransport() ServerTransport
	CreateCodec() Codec
	CreateLogPlugin() LogPlugin
}

// 具体A服务工厂
type AServiceFactory struct {

}

func (a *AServiceFactory) CreateServerTransport() ServerTransport {
	return new(UDPTransport)
}

func (a *AServiceFactory) CreateCodec() Codec {
	return new(QZACodec)
}

func (a *AServiceFactory) CreateLogPlugin() LogPlugin {
	return new(AttaLogger)
}

// 具体B服务工厂
type BServiceFactory struct {

}

func (b *BServiceFactory) CreateServerTransport() ServerTransport {
	return new(TCPTransport)
}

func (b *BServiceFactory) CreateCodec() Codec {
	return new(OIDBCodec)
}

func (b *BServiceFactory) CreateLogPlugin() LogPlugin {
	return new(ConsoleWriteLogger)
}