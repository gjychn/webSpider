package log

import (
	"go.uber.org/zap"
	"time"
)

type logger interface {
	Debug(msg string, args ...interface{})
	Info(mgs string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}


func test() {
	_ = zap.NewDevelopmentConfig()


	logger, err := zap.NewProduction()
	if err != nil {
		return
	}
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "url"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

}