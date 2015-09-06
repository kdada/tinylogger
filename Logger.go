// logger 包实现了一个控制台或文件日志均可用的日志记录器
package logger

import (
	"fmt"
)

//日志记录器
//实现:ILogger
type Logger struct {
	logLevel   LogLevel    //日志输出等级
	logChannel chan string //日志管道
	logWriter  ILogWriter  //日志写入
	cloesd     bool        //日志是否已经关闭
}

// newLogger 创建日志记录器
func newLogger() *Logger {
	var logger = new(Logger)
	logger.logLevel = LogLevelDebug | LogLevelInfo | LogLevelException | LogLevelError
	logger.logChannel = make(chan string, 100)
	return logger
}

// NewConsoleLogger 创建控制台日志记录器
func NewConsoleLogger() *Logger {
	var logger = newLogger()
	var writer, err = NewConsoleLogWriter()
	if err != nil {
		fmt.Println("无法创建控制台日志写入器:" + err.Error())
	}
	logger.logWriter = writer
	logger.logWriter.AsyncWrite(logger.logChannel)
	return logger
}

// NewFileLogger 创建文件日志记录器
func NewFileLogger() *Logger {
	var logger = newLogger()
	var writer, err = NewFileLogWriter()
	if err != nil {
		fmt.Println("无法创建文件日志写入器:" + err.Error())
	}
	logger.logWriter = writer
	logger.logWriter.AsyncWrite(logger.logChannel)
	return logger
}

// LogLevel 得到日志等级是否输出
func (this *Logger) LogLevel(level LogLevel) bool {
	return this.logLevel&level > 0
}

// SetLogLevel 改变某个日志等级是否输出
func (this *Logger) SetLogLevel(level LogLevel, output bool) {
	if output {
		this.logLevel |= level
	} else {
		this.logLevel &= ^level
	}
}

// WriteLog 写入日志
func (this *Logger) WriteLog(level LogLevel, info string) {
	switch level {
	case LogLevelDebug:
		this.WriteDebugLog(info)
	case LogLevelInfo:
		this.WriteInfoLog(info)
	case LogLevelException:
		this.WriteExceptionLog(info)
	case LogLevelError:
		this.WriteErrorLog(info)
	}
}

// WriteDebugLog 写入调试信息
func (this *Logger) WriteDebugLog(info string) {
	if LogLevelDebug&this.logLevel > 0 {
		var log = "[Debug]" + info
		this.logChannel <- log
	}
}

// WriteInfoLog 写入一般信息
func (this *Logger) WriteInfoLog(info string) {
	if LogLevelInfo&this.logLevel > 0 {
		var log = "[Info]" + info
		this.logChannel <- log
	}
}

// WriteExceptionLog 写入异常信息
func (this *Logger) WriteExceptionLog(info string) {
	if LogLevelException&this.logLevel > 0 {
		var log = "[Exception]" + info
		this.logChannel <- log
	}
}

// WriteErrorLog 写入错误信息
func (this *Logger) WriteErrorLog(info string) {
	if LogLevelError&this.logLevel > 0 {
		var log = "[Error]" + info
		this.logChannel <- log
	}
}

// Cloesd 日志是否已关闭
func (this *Logger) Cloesd() bool {
	return this.cloesd
}

// Close 关闭日志 关闭后无法再使用
func (this *Logger) Close() {
	if !this.Cloesd() {
		this.cloesd = true
		this.logWriter.Close()
	}
}
