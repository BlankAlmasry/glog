package awesomeProject2

import "github.com/golang/glog"

type GlogAdapter struct{}

func (a *GlogAdapter) Logf(level LogLevel, message string, args ...interface{}) {
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
	}
}

func (a *GlogAdapter) Log(level LogLevel, message string) {
	switch level {
	case LogLevelInfo:
		if V(LogLevelInfo) {
			glog.Info(message)
		}
	case LogLevelWarn:
		if V(LogLevelWarn) {
			glog.Warning(message)
		}
	case LogLevelError:
		if V(LogLevelError) {
			glog.Error(message)
		}
	case LogLevelFatal:
		if V(LogLevelFatal) {
			glog.Fatal(message)
		}
	}
}

func (a *GlogAdapter) Flush() {
	glog.Flush()
}

/*V
 * glog.V() returns true if the log level is greater than or equal to the specified level (i.e. v >= LogLevel).
 * This function returns true if the log level is less than or equal the specified level (i.e. v <= LogLevel).
 */
func V(level LogLevel) glog.Verbose {
	return !glog.V(glog.Level(level + 1))
}

var logger = &GlogAdapter{}
