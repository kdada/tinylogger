package logger

//日志级别
type LogLevel uint

const (
	LogLevelDebug     LogLevel = 1 << iota //调试信息等级
	LogLevelInfo                           //输出信息等级
	LogLevelException                      //异常信息等级
	LogLevelError                          //错误信息等级
)

// ToString 取得日志等级对应的字符串
func (this LogLevel) ToString() string {
	switch this {
	case LogLevelDebug:
		return "Debug"
	case LogLevelInfo:
		return "Info"
	case LogLevelException:
		return "Exception"
	case LogLevelError:
		return "Error"
	default:
		return "None"
	}
}
