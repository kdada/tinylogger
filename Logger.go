package logger

//日志接口 所有日志类型应该实现该接口
type Logger interface {
	// WriteLog 写入日志
	WriteLog(level LogLevel, info string)
	// WriteDebugLog 写入调试信息
	WriteDebugLog(info string)
	// WriteInfoLog 写入一般信息
	WriteInfoLog(info string)
	// WriteExceptionLog 写入异常信息
	WriteExceptionLog(info string)
	// WriteErrorLog 写入错误信息
	WriteErrorLog(info string)
	// LogLevel 获取某个日志等级是否输出
	LogLevelOutput(level LogLevel) bool
	// SetLogLevelOutput 设置某个日志等级是否输出
	SetLogLevelOutput(level LogLevel, output bool)
	// Async 是否异步输出
	Async() bool
	// SetAsync 设置是否异步输出
	SetAsync(async bool)
	// Close 关闭日志 关闭后无法再进行写入操作
	Close()
	// Cloesd 日志是否关闭
	Cloesd() bool
}
