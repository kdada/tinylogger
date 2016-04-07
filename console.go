package logger

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

//控制台日志写入器
type ConsoleLogWriter struct {
	logList *list.List  // 日志列表
	logmu   *sync.Mutex // 日志列表锁
	closed  bool        // 是否已经停止
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
func (this *ConsoleLogWriter) AsyncWrite(logList *list.List, mu *sync.Mutex) {
	this.logList = logList
	this.logmu = mu
	go func() {
		for !this.closed {
			if this.logList.Len() > 0 {
				var start *list.Element
				var length = 0
				this.logmu.Lock()
				start = this.logList.Front()
				length = this.logList.Len()
				this.logList.Init()
				this.logmu.Unlock()
				for i := 0; i < length; i++ {
					var v, ok = start.Value.(string)
					if ok {
						this.Write(v)
					}
					start = start.Next()
				}
			} else {
				//暂停15毫秒
				time.Sleep(15 * time.Millisecond)
				//runtime.Gosched()
			}
		}
	}()
}

// Close 关闭日志写入器
func (this *ConsoleLogWriter) Close() {
	this.closed = true
}
