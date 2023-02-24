package glog

import "github.com/golang/glog"

func Logf(level LogLevel, message string, args ...interface{}) {
	switch level {
	case LogLevelInfo:
		if V(LogLevelInfo) {
			glog.Infof(message, args...)
		}
	case LogLevelWarn:
		if V(LogLevelWarn) {
			glog.Warningf(message, args...)
		}
	case LogLevelError:
		if V(LogLevelError) {
			glog.Errorf(message, args...)
		}
	case LogLevelFatal:
		if V(LogLevelFatal) {
			glog.Fatalf(message, args...)
		}
	case LogLevelExit:
		if V(LogLevelExit) {
			glog.Exitf(message, args...)
		}
	}
}

func Log(level LogLevel, args ...interface{}) {
	switch level {
	case LogLevelInfo:
		if V(LogLevelInfo) {
			glog.Info(args...)
		}
	case LogLevelWarn:
		if V(LogLevelWarn) {
			glog.Warning(args...)
		}
	case LogLevelError:
		if V(LogLevelError) {
			glog.Error(args...)
		}
	case LogLevelFatal:
		if V(LogLevelFatal) {
			glog.Fatal(args...)
		}
	case LogLevelExit:
		if V(LogLevelExit) {
			glog.Exit(args...)
		}
	}
}

func Logln(level LogLevel, args ...interface{}) {
	switch level {
	case LogLevelInfo:
		if V(LogLevelInfo) {
			glog.Infoln(args...)
		}
	case LogLevelWarn:
		if V(LogLevelWarn) {
			glog.Warningln(args...)
		}
	case LogLevelError:
		if V(LogLevelError) {
			glog.Errorln(args...)
		}
	case LogLevelFatal:
		if V(LogLevelFatal) {
			glog.Fatalln(args...)
		}
	case LogLevelExit:
		if V(LogLevelExit) {
			glog.Exitln(args...)
		}
	}
}

func Flush() {
	glog.Flush()
}

/*V
 * glog.V() returns true if the log level is greater than or equal to the specified level (i.e. v >= LogLevel).
 * This function returns true if the log level is less than or equal the specified level (i.e. v <= LogLevel).
 */
func V(level LogLevel) glog.Verbose {
	return !glog.V(glog.Level(level + 1))
}
