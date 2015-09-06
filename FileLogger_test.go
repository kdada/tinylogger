package logger

import (
	"fmt"
	"testing"
)

var fileLogger ILogger

//测试输出
func TestFileLoggerOutput(t *testing.T) {
	var logger = fileLogger
	var log = "T:测试测试"
	logger.WriteDebugLog(log)
	logger.WriteInfoLog(log)
	logger.WriteExceptionLog(log)
	logger.WriteErrorLog(log)

	log += "type"
	logger.WriteLog(LogLevelDebug, log)
	logger.WriteLog(LogLevelInfo, log)
	logger.WriteLog(LogLevelException, log)
	logger.WriteLog(LogLevelError, log)
}

//性能测试 100goroutine测试
func BenchmarkFileLoggerOutput(b *testing.B) {
	var logger = fileLogger
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		var i = 0
		for pb.Next() {
			i++
			var log = fmt.Sprintf("%d B:测试测试", i)
			logger.WriteDebugLog(log)
			logger.WriteInfoLog(log)
			logger.WriteExceptionLog(log)
			logger.WriteErrorLog(log)

			log += "type"
			logger.WriteLog(LogLevelDebug, log)
			logger.WriteLog(LogLevelInfo, log)
			logger.WriteLog(LogLevelException, log)
			logger.WriteLog(LogLevelError, log)
		}
	})

}
