package glog

type Logger interface {
	Log(level LogLevel, message string, args ...interface{})
}

func Infof(message string, args ...interface{}) {
	Logf(LogLevelInfo, message, args...)
}

func Warningf(message string, args ...interface{}) {
	Logf(LogLevelWarn, message, args...)
}

func Errorf(message string, args ...interface{}) {
	Logf(LogLevelError, message, args...)
}

func Fatalf(message string, args ...interface{}) {
	Logf(LogLevelFatal, message, args...)
}

func Exitf(message string, args ...interface{}) {
	Logf(LogLevelExit, message, args...)
}

func Info(args ...interface{}) {
	Log(LogLevelInfo, args...)
}

func Warning(args ...interface{}) {
	Log(LogLevelWarn, args...)
}

func Error(args ...interface{}) {
	Log(LogLevelError, args...)
}

func Fatal(args ...interface{}) {
	Log(LogLevelFatal, args...)
}

func Exit(args ...interface{}) {
	Log(LogLevelExit, args...)
}

func Infoln(args ...interface{}) {
	Logln(LogLevelInfo, args...)
}

func Warningln(args ...interface{}) {
	Logln(LogLevelWarn, args...)
}

func Errorln(args ...interface{}) {
	Logln(LogLevelError, args...)
}

func Fatalln(args ...interface{}) {
	Logln(LogLevelFatal, args...)
}

func Exitln(args ...interface{}) {
	Logln(LogLevelExit, args...)
}
