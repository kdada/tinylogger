// Package logger 包实现了一个控制台或文件日志均可用的日志记录器
package logger

import "sync"

//日志接口 所有日志类型应该实现该接口
type Logger interface {
	// WriteDebugLog 写入调试信息
	WriteDebugLog(info ...interface{})
	// WriteInfoLog 写入一般信息
	WriteInfoLog(info ...interface{})
	// WriteExceptionLog 写入异常信息
	WriteExceptionLog(info ...interface{})
	// WriteErrorLog 写入错误信息
	WriteErrorLog(info ...interface{})
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

type LoggerCreator func() (Logger, error)

var (
	mu       sync.Mutex                           //生成器互斥锁
	creators = make(map[LoggerType]LoggerCreator) //日志创建器映射
)

// NewLogger 创建一个新的Logger
//  kind:日志类型
func NewLogger(kind LoggerType) (Logger, error) {
	var creator, ok = creators[kind]
	if !ok {
		return nil, LoggerErrorInvalidKind.Format(kind).Error()
	}
	return creator()
}

// RegisterLoggerCreator 注册LoggerCreator创建器
func RegisterLoggerCreator(kind LoggerType, creator LoggerCreator) error {
	mu.Lock()
	defer mu.Unlock()
	creators[kind] = creator
	return nil
}
