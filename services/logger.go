package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

type LogService struct {
	*logrus.Logger
}

type CtxLogger struct {
	*logrus.Entry
}

func NewLogger(level string) (*LogService) {
	logger := LogService{}
	logger.Logger = logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = &prefixed.TextFormatter{}
	switch level {
	case "error":
		logger.Level = logrus.ErrorLevel
	case "info":
		logger.Level = logrus.InfoLevel
	default:
		logger.Level = logrus.DebugLevel
	}
	return &logger
}

func (ls *LogService) NewPrefix(prefix string) *CtxLogger {
	ctxLog := CtxLogger{
		Entry: ls.WithField("prefix", prefix),
	}
	return &ctxLog
}

func (ctxl *CtxLogger) NewPrefix(prefix string) *CtxLogger {
	ctxLog := CtxLogger{
		Entry: ctxl.WithField("prefix", prefix),
	}
	return &ctxLog
}

func (ctxl *CtxLogger) Print(v ...interface{}) {
	ctxl.Debug(v)
}

func (ctxl *CtxLogger) AddPrefix(prefix string) *CtxLogger {
	var newPrefix string
	if data, ok := ctxl.Data["prefix"]; !ok {
		newPrefix = prefix
	} else {
		newPrefix = fmt.Sprintf("%s.%s", data, prefix)
	}
	return &CtxLogger{
		Entry: ctxl.WithField("prefix", newPrefix),
	}
}
