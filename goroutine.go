package logger

import (
	"fmt"
)

// LogGoroutineCallback 定义用于执行的callback
type LogGoroutineCallback func(logger Logger)

// LogGoroutine 执行callback,并在callback中发生panic时自动恢复并使用logger记录出错的信息
func LogGoroutine(logger Logger, callback LogGoroutineCallback) {
	defer func() {
		if err := recover(); err != nil {
			var errInfo = fmt.Sprint(err)
			logger.WriteErrorLog("goroutine崩溃:" + errInfo)
		}
	}()
	callback(logger)
}

// AsyncLogGoroutine 异步执行callback,并在callback中发生panic时自动恢复并使用logger记录出错的信息
func AsyncLogGoroutine(logger Logger, callback LogGoroutineCallback) {
	go LogGoroutine(logger, callback)
}
