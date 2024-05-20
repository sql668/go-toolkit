package log

import (
	"bytes"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestStdLogger(_ *testing.T) {
	logger := DefaultLogger
	logger = With(logger, "caller", DefaultCaller, "ts", DefaultTimestamp)

	_ = logger.Log(DebugLevel, "msg", "test debug")
	_ = logger.Log(InfoLevel, "msg", "test info")
	_ = logger.Log(WarnLevel, "msg", "test warn")
	_ = logger.Log(ErrorLevel, "msg", "test error")
	_ = logger.Log(DebugLevel, "singular")

	logger2 := DefaultLogger
	_ = logger2.Log(DebugLevel)
}

func TestStdLogger_Log(t *testing.T) {
	var b bytes.Buffer
	logger := NewStdLogger(&b)

	var eg errgroup.Group
	eg.Go(func() error { return logger.Log(InfoLevel, "msg", "a", "k", "v") })
	eg.Go(func() error { return logger.Log(InfoLevel, "msg", "a", "k", "v") })

	err := eg.Wait()
	if err != nil {
		t.Fatalf("log error: %v", err)
	}

	if s := b.String(); s != "INFO msg=a k=v\nINFO msg=a k=v\n" {
		t.Fatalf("log not match: %q", s)
	}
}
