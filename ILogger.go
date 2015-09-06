package logger

//日志接口 所有日志类型应该实现该接口
type ILogger interface {
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
	// Close 关闭日志 关闭后无法再进行写入操作
	Close()
}
