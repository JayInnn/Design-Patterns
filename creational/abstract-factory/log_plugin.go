package abstract_factory

type LogPlugin interface {
	WriteLog()
}

// 控制台日志插件
type ConsoleWriteLogger struct {

}

func (c *ConsoleWriteLogger) WriteLog()  {

}

// 文件日志插件
type FileWriteLogger struct {

}

func (f *FileWriteLogger) WriteLog()  {

}

// atta日志插件
type AttaLogger struct {

}

func (a *AttaLogger) WriteLog()  {

}