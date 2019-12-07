package log

import (
	"fmt"
	"go.uber.org/zap"
)

type zapLogger struct {
	 *zap.Logger
}

func NewZapperLogger()  *zapLogger {

	conf := zap.NewDevelopmentConfig()
	logger, _ := conf.Build(zap.AddCaller(), zap.AddCallerSkip(1))

	return &zapLogger{logger}
}

func (z *zapLogger)Debug(msg string, args ...interface{})  {
	defer z.Logger.Sync()

	z.Logger.Debug(fmt.Sprintf(msg, args...))
}

func (z *zapLogger)Info(msg string, args ...interface{})  {
	defer z.Logger.Sync()

	z.Logger.Debug(fmt.Sprintf(msg, args...))
}

func (z *zapLogger)Warn(msg string, args ...interface{})  {
	defer z.Logger.Sync()

	z.Logger.Debug(fmt.Sprintf(msg, args...))
}

func (z *zapLogger)Error(msg string, args ...interface{})  {
	defer z.Logger.Sync()

	z.Logger.Debug(fmt.Sprintf(msg, args...))
}
