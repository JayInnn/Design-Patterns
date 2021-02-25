package factory_method

type LogFactory interface {
	createLogger() Logger
}

// 打印台输出
type ConsoleWriterFactory struct {
}

func (c *ConsoleWriterFactory) createLogger() Logger {
	return new(ConsoleWriterLogger)
}


// 文件输出
type FileWriterFactory struct {
}

func (f *FileWriterFactory) createLogger() Logger {
	return new(FileWriterLogger)
}

