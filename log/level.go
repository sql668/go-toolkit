package log

import (
	"strings"
)

type Level int8

// LevelKey is logger level key.
const LevelKey = "level"

const (
	// TraceLevel 日志级别很低，很少使用
	TraceLevel Level = iota - 2
	// DebugLevel  主要用于开发过程中打印运行信息
	DebugLevel
	// InfoLevel 用于输出程序运行的重要信息或感兴趣的信息，为默认级别
	InfoLevel
	// WarnLevel 表明会出现潜在错误的情形
	WarnLevel
	// ErrorLevel 打印错误和异常信息
	ErrorLevel
	// FatalLevel 重大错误，出现这种级别的事件会导致程序不能继续运行
	FatalLevel
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "TRACE"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	}
	return ""
}

//// Enabled returns true if the given level is at or above this level.
//func (l Level) Enabled(lvl Level) bool {
//	return lvl >= l
//}

func (l Level) Key() string {
	return LevelKey
}


// ParseLevel 从字符串level中得到 logger level
func ParseLevel(levelStr string) (Level) {
	switch strings.ToUpper(levelStr) {
	case TraceLevel.String():
		return TraceLevel
	case DebugLevel.String():
		return DebugLevel
	case InfoLevel.String():
		return InfoLevel
	case WarnLevel.String():
		return WarnLevel
	case ErrorLevel.String():
		return ErrorLevel
	case FatalLevel.String():
		return FatalLevel
	}
	return InfoLevel
}