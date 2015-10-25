package logger

import (
	"fmt"
)

func init() {
	//注册控制台日志创建器
	RegisterLoggerCreator(LoggerTypeConsole, newConsoleLogger)
	//注册文件日志创建器
	RegisterLoggerCreator(LoggerTypeFile, newFileLogger)

	//初始化默认日志记录器
	var err error
	defaultLogger, err = NewLogger(LoggerTypeConsole)
	defaultLogger.SetAsync(false)
	if err != nil {
		fmt.Println(err)
	}
}
