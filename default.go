package logger

//默认日志记录器
var defaultLogger Logger = NewFileLogger()

// DefaultLogger 返回默认日志记录器(默认为文件日志记录器)
func DefaultLogger() Logger {
	return defaultLogger
}

// SetDefaultLogger 设置默认日志记录器
func SetDefaultLogger(logger Logger) {
	defaultLogger = logger
}




