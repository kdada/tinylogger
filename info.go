package logger

// 日志类型
type LoggerType string

const (
	LoggerTypeConsole LoggerType = "console" // 控制台日志类型
	LoggerTypeFile    LoggerType = "file"    //文件日志类型
)

// 日志时间格式
const timeFormat = "2006-01-02 15:04:05.000000"
