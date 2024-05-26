package log

// 有时日志中可能会有敏感信息，需要进行脱敏，或者只打印级别高的日志，这时候就可以使用Filter来对日志的输出进行一些过滤操作，通常用法是使用Filter来包装原始的Logger，用来创建Helper使用。
// FilterLevel 按照日志等级过滤，低于该等级的日志将不会被输出。例如这里传入FilterLevel(logger.ErrorLevel)，则debug/info/warn日志都会被过滤掉不会输出，error和fatal正常输出。
// FilterKey(key ...string) FilterOption 按照key过滤，这些key的值会被***遮蔽
// FilterValue(value ...string) FilterOption 按照value过滤，匹配的值会被***遮蔽
// FilterFunc(f func(level Level, keyvals ...interface{}) bool) 使用自定义的函数来对日志进行处理，keyvals里为key和对应的value，按照奇偶进行读取即可

/**
h := NewHelper(
		NewFilter(logger,
			// 等级过滤
			FilterLevel(log.LevelError),

			// 按key遮蔽
			FilterKey("username"),

			// 按value遮蔽
			FilterValue("hello"),

			// 自定义过滤函数
			FilterFunc(
				func (level Level, keyvals ...interface{}) bool {
					if level == LevelWarn {
						return true
					}
					for i := 0; i < len(keyvals); i++ {
						if keyvals[i] == "password" {
							keyvals[i+1] = fuzzyStr
						}
					}
					return false
				}
			),
		),
	)

	h.Log(log.LevelDebug, "msg", "test debug")
	h.Info("hello")
	h.Infow("password", "123456")
	h.Infow("username", "kratos")
	h.Warn("warn log")
*/
// FilterOption is filter option.
type FilterOption func(*Filter)

const fuzzyStr = "***"

// FilterLevel 按照日志等级过滤，低于该等级的日志将不会被输出
func FilterLevel(level Level) FilterOption {
	return func(opts *Filter) {
		opts.level = level
	}
}

// FilterKey 按照key过滤，这些key的值会被***遮蔽
func FilterKey(key ...string) FilterOption {
	return func(o *Filter) {
		for _, v := range key {
			o.key[v] = struct{}{}
		}
	}
}

// FilterValue 按照value过滤，匹配的值会被***遮蔽
func FilterValue(value ...string) FilterOption {
	return func(o *Filter) {
		for _, v := range value {
			o.value[v] = struct{}{}
		}
	}
}

// FilterFunc 使用自定义的函数来对日志进行处理，keyvals里为key和对应的value，按照奇偶进行读取即可
func FilterFunc(f func(level Level, keyvals ...interface{}) bool) FilterOption {
	return func(o *Filter) {
		o.filter = f
	}
}

// Filter is a logger filter.
type Filter struct {
	logger Logger
	level  Level
	key    map[interface{}]struct{}
	value  map[interface{}]struct{}
	filter func(level Level, keyvals ...interface{}) bool
}

// NewFilter new a logger filter.
func NewFilter(logger Logger, opts ...FilterOption) *Filter {
	options := Filter{
		logger: logger,
		key:    make(map[interface{}]struct{}),
		value:  make(map[interface{}]struct{}),
	}
	for _, o := range opts {
		o(&options)
	}
	return &options
}

// Log Print log by level and keyvals.
func (f *Filter) Log(level Level, keyvals ...interface{}) error {
	if level < f.level {
		return nil
	}
	// prefixkv contains the slice of arguments defined as prefixes during the log initialization
	var prefixkv []interface{}
	l, ok := f.logger.(*logger)
	if ok && len(l.prefix) > 0 {
		prefixkv = make([]interface{}, 0, len(l.prefix))
		prefixkv = append(prefixkv, l.prefix...)
	}

	if f.filter != nil && (f.filter(level, prefixkv...) || f.filter(level, keyvals...)) {
		return nil
	}

	if len(f.key) > 0 || len(f.value) > 0 {
		for i := 0; i < len(keyvals); i += 2 {
			v := i + 1
			if v >= len(keyvals) {
				continue
			}
			if _, ok := f.key[keyvals[i]]; ok {
				keyvals[v] = fuzzyStr
			}
			if _, ok := f.value[keyvals[v]]; ok {
				keyvals[v] = fuzzyStr
			}
		}
	}
	return f.logger.Log(level, keyvals...)
}
