package infra

import (
	"context"
	"github.com/labstack/echo/log"
	log2 "github.com/labstack/gommon/log"

	"io"
)

func NewAppEngineLogger(ctx context.Context) log.Logger {
	return &AppEngineLogger{ctx}
}

type AppEngineLogger struct {
	ctx context.Context
}

func (l *AppEngineLogger) SetOutput(io.Writer) {
	panic("implement me")
}

func (l *AppEngineLogger) SetLevel(log2.Lvl) {
	panic("implement me")
}

func (l *AppEngineLogger) Print(...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Printf(string, ...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Printj(log2.JSON) {
	panic("implement me")
}

func (l *AppEngineLogger) Debug(...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Debugf(string, ...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Debugj(log2.JSON) {
	panic("implement me")
}

func (l *AppEngineLogger) Info(...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Infof(string, ...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Infoj(log2.JSON) {
	panic("implement me")
}

func (l *AppEngineLogger) Warn(...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Warnf(string, ...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Warnj(log2.JSON) {
	panic("implement me")
}

func (l *AppEngineLogger) Error(...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Errorf(string, ...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Errorj(log2.JSON) {
	panic("implement me")
}

func (l *AppEngineLogger) Fatal(...interface{}) {
	panic("implement me")
}

func (l *AppEngineLogger) Fatalj(log2.JSON) {
	panic("implement me")
}

func (l *AppEngineLogger) Fatalf(string, ...interface{}) {
	panic("implement me")
}


