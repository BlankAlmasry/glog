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

func Exitf(message string, args ...interface{}) {
	logger.Logf(LogLevelExit, message, args...)
}

func Info(args ...interface{}) {
	logger.Log(LogLevelInfo, args)
}

func Warning(args ...interface{}) {
	logger.Log(LogLevelWarn, args)
}

func Error(args ...interface{}) {
	logger.Log(LogLevelError, args)
}

func Fatal(args ...interface{}) {
	logger.Log(LogLevelFatal, args)
}

func Exit(args ...interface{}) {
	logger.Log(LogLevelExit, args)
}

func Infoln(args ...interface{}) {
	logger.Logln(LogLevelInfo, args...)
}

func Warningln(args ...interface{}) {
	logger.Logln(LogLevelWarn, args...)
}

func Errorln(args ...interface{}) {
	logger.Logln(LogLevelError, args...)
}

func Fatalln(args ...interface{}) {
	logger.Logln(LogLevelFatal, args...)
}

func Exitln(args ...interface{}) {
	logger.Logln(LogLevelExit, args...)
}

func Flush() {
	logger.Flush()
}
