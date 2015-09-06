package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

//文件日志写入器
//实现:ILogWriter
type FileLogWriter struct {
	file   *os.File      //日志文件
	writer *bufio.Writer //写入工具
	day    int           //文件日期
}

// NewFileLogWriter 创建文件日志写入器
func NewFileLogWriter() (*FileLogWriter, error) {
	//创建日志目录
	var err = os.Mkdir("logs", 0770)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	//创建写入器
	var writer = new(FileLogWriter)
	writer.day = 0
	err = writer.createLogFile(time.Now())
	if err != nil {
		return nil, err
	}
	return writer, nil
}

// createLogFile 创建日志文件
func (this *FileLogWriter) createLogFile(date time.Time) error {
	var day = date.Day()
	if day == this.day {
		//文件无需更新
		return nil
	}
	//关闭原来的日志文件,并创建新的日志文件
	if this.file != nil {
		err := this.file.Close()
		if err != nil {
			return err
		}
	}
	//创建新的日志文件
	var fileName = date.Format("2006-01-02") + ".log"
	file, err := os.OpenFile("logs/"+fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil && !os.IsExist(err) {
		return err
	}
	this.file = file
	this.writer = bufio.NewWriter(file)
	return nil
}

// Write 写入日志
func (this *FileLogWriter) Write(log string) {
	var date = time.Now()
	var err = this.createLogFile(date)
	if err == nil {
		this.writer.WriteString(date.Format("2006-01-02 15:04:05 ") + log + "\n")
		this.writer.Flush()
	} else {
		fmt.Println("写入日志出错:" + err.Error())
	}
}

// AsyncWrite 异步写入日志
func (this *FileLogWriter) AsyncWrite(logChannel chan string) {
	go func() {
		for true {
			this.Write(<-logChannel)
		}
	}()
}

// Close 关闭日志写入器
func (this *FileLogWriter) Close() {
	this.file.Close()
}
