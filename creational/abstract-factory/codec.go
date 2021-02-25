package abstract_factory

type Codec interface {
	Encode()
	Decode()
}

// Oidb协议编解码器
type OIDBCodec struct {

}

func (o *OIDBCodec) Encode() {

}

func (o *OIDBCodec) Decode() {

}


// qza协议编解码器
type QZACodec struct {

}

func (q *QZACodec) Encode() {

}

func (q *QZACodec) Decode() {

}