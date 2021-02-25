package factory_method

type Logger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

// 打印台输出
type ConsoleWriterLogger struct {

}

func (* ConsoleWriterLogger) Error(args ...interface{}) {

}

func (* ConsoleWriterLogger) Errorf(format string, args ...interface{}) {

}

// 文件输出
type FileWriterLogger struct {

}

func (* FileWriterLogger) Error(args ...interface{}) {

}

func (* FileWriterLogger) Errorf(format string, args ...interface{}) {

}