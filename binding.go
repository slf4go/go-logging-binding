package go_logging_connector

import (
	goLogging "github.com/op/go-logging"
	slf4go "github.com/slf4go/logger"
)

const module = "slf4go"

var log = goLogging.MustGetLogger(module)
var levelMap = map[slf4go.Level]goLogging.Level{
	slf4go.LogPanic:  goLogging.CRITICAL,
	slf4go.LogError:  goLogging.ERROR,
	slf4go.LogWarn:   goLogging.WARNING,
	slf4go.LogNotice: goLogging.NOTICE,
	slf4go.LogInfo:   goLogging.INFO,
	slf4go.LogDebug:  goLogging.DEBUG,
	slf4go.LogTrace:  goLogging.DEBUG,
}

func init() {
	slf4go.BindLogImpl(GoLoggingImpl{})
}

type GoLoggingImpl struct{}

func (GoLoggingImpl) SetLevel(level slf4go.Level) {
	goLogging.SetLevel(levelMap[level], module)
}

func (GoLoggingImpl) Log(level slf4go.Level, msg string, stack []string) {
	var logFunc func(...interface{})

	switch level {
	case slf4go.LogPanic:
		logFunc = log.Critical
	case slf4go.LogError:
		logFunc = log.Error
	case slf4go.LogWarn:
		logFunc = log.Warning
	case slf4go.LogNotice:
		logFunc = log.Notice
	case slf4go.LogInfo:
		logFunc = log.Info
	case slf4go.LogDebug:
		logFunc = log.Debug
	}

	logFunc(msg)
	if stack != nil {
		for i, line := range stack {
			logFunc("%d: %s", len(stack)-i-1, line)
		}
	}
}
