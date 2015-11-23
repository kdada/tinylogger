# tinylogger
日志记录器

1.控制台日志记录  
2.文件日志记录(按日期区分文件,每天创建一个日志文件)  

DefaultLogger是默认创建的控制台日志记录器,并且是同步的,可以通过SetAsync设置异步  

```go
package main

import (
	"time"

	"github.com/kdada/tinylogger"
)

func main() {
	//logger包默认日志使用的是控制台同步日志
	logger.DefaultLogger().WriteDebugLog("控制台日志测试")
	//创建一个文件异步日志
	var l = logger.NewFileLogWriter()
	l.WriteDebugLog("文件日志测试")
	//文件日志默认为异步模式,因此需要在主线程中等待一会
	time.Sleep(1 * time.Second)
}
```


输出结果如下:  
在main程序的同级目录中会创建一个logs目录  
生成文件2015-10-15.log  
文件内容如下格式(日期.微秒[日志类型]日志内容):  
2015-10-15 17:53:30.209630[Debug]文件日志测试  


控制台输出如下:  
2015-10-15 17:53:30.209630[Debug]控制台日志测试  
