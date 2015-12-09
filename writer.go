package logger

import (
	"container/list"
	"sync"
)

//日志写入类型接口
type LogWriter interface {
	// Write 日志写入
	Write(log string)
	// AsyncWrite 日志持续写入,数据源来自logChannel
	AsyncWrite(logList *list.List, mu *sync.Mutex)
	// Close 关闭写入器
	Close()
}
