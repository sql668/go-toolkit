package log

import (
	"context"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	// DefaultCaller 打印出调用日志方法文件名和行号，例如 caller=logger/log_test.go:11
	DefaultCaller = Caller(3)

	// DefaultTimestamp 打印时间戳
	DefaultTimestamp = Timestamp(time.RFC3339)
)


type Valuer func(ctx context.Context) interface{}

// Value return the function value.
func Value(ctx context.Context, v interface{}) interface{} {
	if v, ok := v.(Valuer); ok {
		return v(ctx)
	}
	return v
}

// Caller returns a Valuer that returns a pkg/file:line description of the caller.
func Caller(depth int) Valuer {
	return func(context.Context) interface{} {
		_, file, line, _ := runtime.Caller(depth)
		idx := strings.LastIndexByte(file, '/')
		if idx == -1 {
			return file[idx+1:] + ":" + strconv.Itoa(line)
		}
		idx = strings.LastIndexByte(file[:idx], '/')
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

// Timestamp returns a timestamp Valuer with a custom time format.
func Timestamp(layout string) Valuer {
	return func(context.Context) interface{} {
		return time.Now().Format(layout)
	}
}

func
bindValues(ctx context.Context, keyAndVals []interface{}) {
	for i := 1; i < len(keyAndVals); i += 2 {
		if v, ok := keyAndVals[i].(Valuer); ok {
			keyAndVals[i] = v(ctx)
		}
	}
}

func containsValuer(keyAndVals []interface{}) bool {
	for i := 1; i < len(keyAndVals); i += 2 {
		if _, ok := keyAndVals[i].(Valuer); ok {
			return true
		}
	}
	return false
}
