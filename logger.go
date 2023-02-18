package awesomeProject2

type Logger interface {
	Log(level LogLevel, message string, args ...interface{})
}

func Infof(message string, args ...interface{}) {
	logger.Logf(LogLevelInfo, message, args...)
}

func Warningf(message string, args ...interface{}) {
	logger.Logf(LogLevelWarn, message, args...)
}

func Errorf(message string, args ...interface{}) {
	logger.Logf(LogLevelError, message, args...)
}

func Fatalf(message string, args ...interface{}) {
	logger.Logf(LogLevelFatal, message, args...)
}

func Info(message string) {
	logger.Log(LogLevelInfo, message)
}

func Warning(message string) {
	logger.Log(LogLevelWarn, message)
}

func Error(message string) {
	logger.Log(LogLevelError, message)
}

func Fatal(message string) {
	logger.Log(LogLevelFatal, message)
}

func Flush() {
	logger.Flush()
}
