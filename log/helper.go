package log

import (
	"context"
	"fmt"
	"os"
)

// DefaultMessageKey default message key.
var DefaultMessageKey = "msg"

// Option is Helper option.
type Option func(*Helper)

// Helper is a logger helper.
type Helper struct {
	logger  Logger
	msgKey  string
	sprint  func(...interface{}) string  // 一般输出
	sprintf func(format string, a ...interface{}) string  // 格式化输出
}

// WithMessageKey with message key.
func WithMessageKey(k string) Option {
	return func(opts *Helper) {
		opts.msgKey = k
	}
}

// WithSprint 自定义输出函数
func WithSprint(sprint func(...interface{}) string) Option {
	return func(opts *Helper) {
		opts.sprint = sprint
	}
}

// WithSprintf 自定义格式化输出函数
func WithSprintf(sprintf func(format string, a ...interface{}) string) Option {
	return func(opts *Helper) {
		opts.sprintf = sprintf
	}
}

// NewHelper new a logger helper.
func NewHelper(logger Logger, opts ...Option) *Helper {
	options := &Helper{
		msgKey:  DefaultMessageKey, // default message key
		logger:  logger,
		sprint:  fmt.Sprint,
		sprintf: fmt.Sprintf,
	}
	for _, o := range opts {
		o(options)
	}
	return options
}

// WithContext returns a shallow copy of h with its context changed
// to ctx. The provided ctx must be non-nil.
func (h *Helper) WithContext(ctx context.Context) *Helper {
	return &Helper{
		msgKey:  h.msgKey,
		logger:  WithContext(ctx, h.logger),
		sprint:  h.sprint,
		sprintf: h.sprintf,
	}
}

// Enabled returns true if the given level above this level.
// It delegates to the underlying *Filter.
func (h *Helper) Enabled(level Level) bool {
	if l, ok := h.logger.(*Filter); ok {
		return level >= l.level
	}
	return true
}

// Log Print log by level and keyAndVals.
func (h *Helper) Log(level Level, keyAndVals ...interface{}) {
	_ = h.logger.Log(level, keyAndVals...)
}

// Trace logs a message at trace level.
func (h *Helper) Trace(a ...interface{}) {
	if !h.Enabled(TraceLevel) {
		return
	}
	_ = h.logger.Log(TraceLevel, h.msgKey, h.sprint(a...))
}

// Tracef logs a message at trace level.
func (h *Helper) Tracef(format string, a ...interface{}) {
	if !h.Enabled(TraceLevel) {
		return
	}
	_ = h.logger.Log(TraceLevel, h.msgKey, h.sprintf(format, a...))
}

// Tracew logs a message at trace level.
func (h *Helper) Tracew(keyAndVals ...interface{}) {
	_ = h.logger.Log(TraceLevel, keyAndVals...)
}

// Debug logs a message at debug level.
func (h *Helper) Debug(a ...interface{}) {
	if !h.Enabled(DebugLevel) {
		return
	}
	_ = h.logger.Log(DebugLevel, h.msgKey, h.sprint(a...))
}

// Debugf logs a message at debug level.
func (h *Helper) Debugf(format string, a ...interface{}) {
	if !h.Enabled(DebugLevel) {
		return
	}
	_ = h.logger.Log(DebugLevel, h.msgKey, h.sprintf(format, a...))
}

// Debugw logs a message at debug level.
func (h *Helper) Debugw(keyAndVals ...interface{}) {
	_ = h.logger.Log(DebugLevel, keyAndVals...)
}

// Info logs a message at info level.
func (h *Helper) Info(a ...interface{}) {
	if !h.Enabled(InfoLevel) {
		return
	}
	_ = h.logger.Log(InfoLevel, h.msgKey, h.sprint(a...))
}

// Infof logs a message at info level.
func (h *Helper) Infof(format string, a ...interface{}) {
	if !h.Enabled(InfoLevel) {
		return
	}
	_ = h.logger.Log(InfoLevel, h.msgKey, h.sprintf(format, a...))
}

// Infow logs a message at info level.
func (h *Helper) Infow(keyAndVals ...interface{}) {
	_ = h.logger.Log(InfoLevel, keyAndVals...)
}

// Warn logs a message at warn level.
func (h *Helper) Warn(a ...interface{}) {
	if !h.Enabled(WarnLevel) {
		return
	}
	_ = h.logger.Log(WarnLevel, h.msgKey, h.sprint(a...))
}

// Warnf logs a message at warnf level.
func (h *Helper) Warnf(format string, a ...interface{}) {
	if !h.Enabled(WarnLevel) {
		return
	}
	_ = h.logger.Log(WarnLevel, h.msgKey, h.sprintf(format, a...))
}

// Warnw logs a message at warnf level.
func (h *Helper) Warnw(keyAndVals ...interface{}) {
	_ = h.logger.Log(WarnLevel, keyAndVals...)
}

// Error logs a message at error level.
func (h *Helper) Error(a ...interface{}) {
	if !h.Enabled(ErrorLevel) {
		return
	}
	_ = h.logger.Log(ErrorLevel, h.msgKey, h.sprint(a...))
}

// Errorf logs a message at error level.
func (h *Helper) Errorf(format string, a ...interface{}) {
	if !h.Enabled(ErrorLevel) {
		return
	}
	_ = h.logger.Log(ErrorLevel, h.msgKey, h.sprintf(format, a...))
}

// Errorw logs a message at error level.
func (h *Helper) Errorw(keyAndVals ...interface{}) {
	_ = h.logger.Log(ErrorLevel, keyAndVals...)
}

// Fatal logs a message at fatal level.
func (h *Helper) Fatal(a ...interface{}) {
	_ = h.logger.Log(FatalLevel, h.msgKey, h.sprint(a...))
	os.Exit(1)
}

// Fatalf logs a message at fatal level.
func (h *Helper) Fatalf(format string, a ...interface{}) {
	_ = h.logger.Log(FatalLevel, h.msgKey, h.sprintf(format, a...))
	os.Exit(1)
}

// Fatalw logs a message at fatal level.
func (h *Helper) Fatalw(keyAndVals ...interface{}) {
	_ = h.logger.Log(FatalLevel, keyAndVals...)
	os.Exit(1)
}
