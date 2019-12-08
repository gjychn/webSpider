package log

type logger interface {
	Debug(msg string, args ...interface{})

	Info(mgs string, args ...interface{})

	Warn(msg string, args ...interface{})

	Error(msg string, args ...interface{})
}
