package decorator

type FileInputSteam struct {
	fd string
}

func (f *FileInputSteam) read() {

}

func (f *FileInputSteam) close() {

}

type ByteArrayInputSteam struct {
	byte []byte
}

func (b *ByteArrayInputSteam) read() {

}

func (b *ByteArrayInputSteam) close() {

}

type StringBufferInputStream struct {
	buff string
}

func (sb *StringBufferInputStream) read() {

}

func (sb *StringBufferInputStream) close() {

}
