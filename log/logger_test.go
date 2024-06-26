package log

import (
	"testing"
)

func TestInfo(_ *testing.T) {
	logger := DefaultLogger
	logger = With(logger, "ts", DefaultTimestamp)
	logger = With(logger, "caller", DefaultCaller)
	_ = logger.Log(InfoLevel, "key1", "value1")
}
