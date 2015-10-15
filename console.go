package logger

import "fmt"

//控制台日志写入器
type ConsoleLogWriter struct {
}

// NewConsoleLogWriter 创建控制台日志写入器
func NewConsoleLogWriter() (*ConsoleLogWriter, error) {
	return new(ConsoleLogWriter), nil
}

// Write 写入日志
func (this *ConsoleLogWriter) Write(log string) {
	fmt.Println(log)
}

// AsyncWrite 异步写入日志
func (this *ConsoleLogWriter) AsyncWrite(logChannel chan string) {
	go func() {
		for true {
			this.Write(<-logChannel)
		}
	}()
}

// Close 关闭日志写入器
func (this *ConsoleLogWriter) Close() {

}
