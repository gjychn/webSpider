package log

import (
	"testing"
)

func TestZapLogger(t *testing.T) {
	NewZapperLogger().Debug("log debug can user")
	NewZapperLogger().Info("log info can user")
	NewZapperLogger().Warn("log warn can user")
	NewZapperLogger().Error("log error can user")
}
