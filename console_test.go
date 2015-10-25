package logger

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var consoleLogger Logger

func TestMain(m *testing.M) {
	var err error
	consoleLogger, err = NewLogger(LoggerTypeConsole)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileLogger, err = NewLogger(LoggerTypeFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Exit(m.Run())
}

//测试输出
func TestConsoleLoggerOutput(t *testing.T) {
	var logger = consoleLogger
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
	time.Sleep(1 * time.Second)
}

//性能测试
func BenchmarkConsoleLoggerOutput(b *testing.B) {
	var logger = consoleLogger
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var log = "B:测试测试"
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
