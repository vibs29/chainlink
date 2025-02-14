package logger

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"go.uber.org/zap/zaptest/observer"

	"github.com/smartcontractkit/chainlink/core/config/envvar"
)

// TestLogger creates a logger that directs output to PrettyConsole configured
// for test output, and to the buffer testMemoryLog. t is optional.
// Log level is derived from the LOG_LEVEL env var.
func TestLogger(tb testing.TB) SugaredLogger {
	return testLogger(tb, nil)
}

// TestLoggerObserved creates a logger with an observer that can be used to
// test emitted logs at the given level or above
func TestLoggerObserved(tb testing.TB, lvl zapcore.Level) (Logger, *observer.ObservedLogs) {
	observedZapCore, observedLogs := observer.New(lvl)
	return testLogger(tb, observedZapCore), observedLogs
}

// testLogger returns a new SugaredLogger for tests. core is optional.
func testLogger(tb testing.TB, core zapcore.Core) SugaredLogger {
	ll, invalid := envvar.LogLevel.Parse()
	a := zap.NewAtomicLevelAt(ll)
	opts := []zaptest.LoggerOption{zaptest.Level(a)}
	if core != nil {
		opts = append(opts, zaptest.WrapOptions(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
			return zapcore.NewTee(c, core)
		})))
	}
	l := &zapLogger{
		level:         a,
		SugaredLogger: zaptest.NewLogger(tb, opts...).Sugar(),
	}
	if invalid != "" {
		l.Error(invalid)
	}
	if tb == nil {
		return Sugared(l)
	}
	return Sugared(l.Named(verShaNameStatic()).Named(tb.Name()))
}
