package logger

import (
	"fmt"
	"testing"
	"time"
)

var fileLogger Logger

//测试输出
func TestFileLoggerOutput(t *testing.T) {
	var logger = fileLogger
	var log = "T:测试测试"
	logger.WriteDebugLog(log)
	logger.WriteInfoLog(log)
	logger.WriteExceptionLog(log)
	logger.WriteErrorLog(log)

	time.Sleep(1 * time.Second)
}

//性能测试
func BenchmarkFileLoggerOutput(b *testing.B) {
	var logger = fileLogger
	b.RunParallel(func(pb *testing.PB) {
		var i = 0
		for pb.Next() {
			i++
			var log = fmt.Sprintf("%d B:测试测试", i)
			logger.WriteDebugLog(log)
			logger.WriteInfoLog(log)
			logger.WriteExceptionLog(log)
			logger.WriteErrorLog(log)

		}
	})
}
