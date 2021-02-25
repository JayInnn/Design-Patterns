package decorator

type BufferedInputStream struct {
	InputSteam
	byte []byte
}

func (b *BufferedInputStream) read() {

}

func (b *BufferedInputStream) close() {

}

func (b *BufferedInputStream) fill() {

}


type LineNumberInputStream struct {
	InputSteam
	lineNumber int
}

func (l *LineNumberInputStream) read() {

}

func (l *LineNumberInputStream) close() {

}

func (l *LineNumberInputStream) GetLineNumber() int {
	return l.lineNumber
}

// 具体使用
var buff = &BufferedInputStream{
	InputSteam: &FileInputSteam{
		fd: "path",
	},
	byte: []byte{0, 1, 2},
}

//type BufferedInputStream struct {
//	FilterInputStream
//	byte []byte
//}
//
//func (b *BufferedInputStream) read() {
//
//}
//
//func (b *BufferedInputStream) close() {
//
//}
//
//func (b *BufferedInputStream) fill() {
//
//}
//
//
//type LineNumberInputStream struct {
//	FilterInputStream
//	lineNumber int
//}
//
//func (l *LineNumberInputStream) read() {
//
//}
//
//func (l *LineNumberInputStream) close() {
//
//}
//
//func (l *LineNumberInputStream) GetLineNumber() int {
//	return l.lineNumber
//}
//
//var buff = &BufferedInputStream{
//	FilterInputStream: FilterInputStream{
//		InputSteam: &FileInputSteam{
//			fd: "path",
//		},
//	},
//	byte: []byte{0, 1, 2},
//}