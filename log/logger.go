package log

import (
	"context"
	"log"
)

var DefaultLogger = NewStdLogger(log.Writer())

type Logger interface {
	Log(level Level, keyAndVals ...interface{}) error // keyAndVals 是一个平铺的键值数组，它的长度需要是偶数，奇数位上的是key，偶数位上的是value
}

type logger struct {
	logger    Logger
	prefix    []interface{}
	hasValuer bool
	ctx       context.Context
}

func (c *logger) Log(level Level, keyAndVals ...interface{}) error {
	kvs := make([]interface{}, 0, len(c.prefix)+len(keyAndVals))
	kvs = append(kvs, c.prefix...)
	if c.hasValuer {
		bindValues(c.ctx, kvs)
	}
	kvs = append(kvs, keyAndVals...)
	return c.logger.Log(level, kvs...)
}


// With logger.With方法会返回一个新的Logger，把参数的Valuer绑上去。
// logger = logger.With(logger, "ts", logger.DefaultTimestamp, "caller", logger.DefaultCaller)
func With(l Logger, kv ...interface{}) Logger {
	c, ok := l.(*logger)
	if !ok {
		return &logger{logger: l, prefix: kv, hasValuer: containsValuer(kv), ctx: context.Background()}
	}
	kvs := make([]interface{}, 0, len(c.prefix)+len(kv))
	kvs = append(kvs, c.prefix...)
	kvs = append(kvs, kv...)
	return &logger{
		logger:    c.logger,
		prefix:    kvs,
		hasValuer: containsValuer(kvs),
		ctx:       c.ctx,
	}
}

// WithContext returns a shallow copy of l with its context changed
// to ctx. The provided ctx must be non-nil.
func WithContext(ctx context.Context, l Logger) Logger {
	switch v := l.(type) {
	default:
		return &logger{logger: l, ctx: ctx}
	case *logger:
		lv := *v
		lv.ctx = ctx
		return &lv
	case *Filter:
		fv := *v
		fv.logger = WithContext(ctx, fv.logger)
		return &fv
	}
}