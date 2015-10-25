package logger

import (
	"fmt"
	"time"
)

//日志记录器
// 默认开启异步模式
// 异步模式:完全不阻塞调用线程
// 同步模式:当缓存通道中日志记录不超过100条时,不阻塞调用线程,超过则阻塞
type logger struct {
	logLevel   LogLevel    //日志输出等级
	logChannel chan string //日志管道
	logWriter  LogWriter   //日志写入
	cloesd     bool        //日志是否已经关闭
	async      bool        //日志是否异步输出
}

// newLogger 创建日志记录器
func newLogger() *logger {
	var logger = new(logger)
	logger.logLevel = LogLevelDebug | LogLevelInfo | LogLevelException | LogLevelError
	logger.logChannel = make(chan string, 100)
	logger.cloesd = false
	logger.async = true
	return logger
}

// newConsoleLogger 创建控制台日志记录器
func newConsoleLogger() (Logger, error) {
	var logger = newLogger()
	var writer, err = NewConsoleLogWriter()
	if err != nil {
		fmt.Println("无法创建控制台日志写入器:" + err.Error())
	}
	logger.logWriter = writer
	logger.logWriter.AsyncWrite(logger.logChannel)
	return logger, nil
}

// newFileLogger 创建文件日志记录器
func newFileLogger() (Logger, error) {
	var logger = newLogger()
	var writer, err = NewFileLogWriter()
	if err != nil {
		fmt.Println("无法创建文件日志写入器:" + err.Error())
	}
	logger.logWriter = writer
	logger.logWriter.AsyncWrite(logger.logChannel)
	return logger, nil
}

// LogLevel 得到日志等级是否输出
func (this *logger) LogLevelOutput(level LogLevel) bool {
	return this.logLevel&level > 0
}

// SetLogLevel 改变某个日志等级是否输出
func (this *logger) SetLogLevelOutput(level LogLevel, output bool) {
	if output {
		this.logLevel |= level
	} else {
		this.logLevel &= ^level
	}
}

// Async 是否异步输出
func (this *logger) Async() bool {
	return this.async
}

// SetAsync 设置是否异步输出
func (this *logger) SetAsync(async bool) {
	this.async = async
}

// WriteLog 写入日志
func (this *logger) WriteLog(level LogLevel, info string) {
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
func (this *logger) WriteDebugLog(info string) {
	if LogLevelDebug&this.logLevel > 0 {
		var log = time.Now().Format("2006-01-02 15:04:05.000000") + "[Debug]" + info
		if this.async {
			go func() {
				this.logChannel <- log
			}()
		} else {
			this.logChannel <- log
		}
	}
}

// WriteInfoLog 写入一般信息
func (this *logger) WriteInfoLog(info string) {
	if LogLevelInfo&this.logLevel > 0 {
		var log = time.Now().Format("2006-01-02 15:04:05.000000") + "[Info]" + info
		if this.async {
			go func() {
				this.logChannel <- log
			}()
		} else {
			this.logChannel <- log
		}
	}
}

// WriteExceptionLog 写入异常信息
func (this *logger) WriteExceptionLog(info string) {
	if LogLevelException&this.logLevel > 0 {
		var log = time.Now().Format("2006-01-02 15:04:05.000000") + "[Exception]" + info
		if this.async {
			go func() {
				this.logChannel <- log
			}()
		} else {
			this.logChannel <- log
		}
	}
}

// WriteErrorLog 写入错误信息
func (this *logger) WriteErrorLog(info string) {
	if LogLevelError&this.logLevel > 0 {
		var log = time.Now().Format("2006-01-02 15:04:05.000000") + "[Error]" + info
		if this.async {
			go func() {
				this.logChannel <- log
			}()
		} else {
			this.logChannel <- log
		}
	}
}

// Cloesd 日志是否已关闭
func (this *logger) Cloesd() bool {
	return this.cloesd
}

// Close 关闭日志 关闭后无法再使用
func (this *logger) Close() {
	if !this.Cloesd() {
		this.cloesd = true
		this.logWriter.Close()
	}
}
