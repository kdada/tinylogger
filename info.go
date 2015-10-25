package logger

// 日志类型
type LoggerType string

const (
	LoggerTypeConsole LoggerType = "console" // 控制台日志类型
	LoggerTypeFile    LoggerType = "file"    //文件日志类型
)
