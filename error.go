package logger

import (
	"errors"
	"fmt"
)

// 配置错误信息
type LoggerError string

// 错误码
const (
	LoggerErrorInvalidKind LoggerError = "L10010:LoggerErrorInvalidKind,无效的Kind(%s)"
)

// Format 格式化错误信息并生成新的错误信息
func (this LoggerError) Format(data ...interface{}) LoggerError {
	return LoggerError(fmt.Sprintf(string(this), data...))
}

// Error 生成error类型
func (this LoggerError) Error() error {
	return errors.New(string(this))
}
