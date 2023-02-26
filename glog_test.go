package glog_test

import (
	"bytes"
	"flag"
	"github.com/blankalmasry/glog"
	"io"
	"os"
	"testing"
)

func captureStderr(f func()) (string, error) {
	flag.Parse()
	err := flag.Set("logtostderr", "true")
	if err != nil {
		return "", err
	}
	old := os.Stderr // keep backup of the real stderr
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}
	os.Stderr = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		if err != nil {
			return
		}
		outC <- buf.String()
	}()

	// calling function which stderr we are going to capture:
	f()

	// back to normal state
	err = w.Close()
	if err != nil {
		return "", err
	}
	os.Stderr = old // restoring the real stderr
	return <-outC, nil
}

func TestLog(t *testing.T) {
	testCases := []struct {
		message          string
		logLevel         glog.LogLevel
		expectedLogLevel string
	}{
		{
			message:          "test info",
			logLevel:         glog.LogLevelInfo,
			expectedLogLevel: "I0224",
		},
		{
			message:          "test warn",
			logLevel:         glog.LogLevelWarn,
			expectedLogLevel: "W0224",
		},
		{
			message:          "test error",
			logLevel:         glog.LogLevelError,
			expectedLogLevel: "E0224",
		},
	}

	for _, testCase := range testCases {
		stdErr, err := captureStderr(func() {
			glog.Log(testCase.logLevel, testCase.message)
		})
		if err != nil {
			t.Errorf("error capturing stderr: %v", err)
		}
		if !bytes.Contains([]byte(stdErr), []byte(testCase.message)) {
			t.Errorf("expected stderr to contain %q, but got %q", testCase.message, stdErr)
		}
		if !bytes.Contains([]byte(stdErr), []byte(testCase.expectedLogLevel)) {
			t.Errorf("expected stderr to contain %q, but got %q", testCase.expectedLogLevel, stdErr)
		}
	}
}

func TestLogf(t *testing.T) {
	testCases := []struct {
		message          string
		logLevel         glog.LogLevel
		expectedLogLevel string
	}{
		{
			message:          "test infof",
			logLevel:         glog.LogLevelInfo,
			expectedLogLevel: "I0224",
		},
		{
			message:          "test warnf",
			logLevel:         glog.LogLevelWarn,
			expectedLogLevel: "W0224",
		},
		{
			message:          "test errorf",
			logLevel:         glog.LogLevelError,
			expectedLogLevel: "E0224",
		},
	}

	for _, testCase := range testCases {
		stdErr, err := captureStderr(func() {
			glog.Logf(testCase.logLevel, testCase.message)
		})
		if err != nil {
			t.Errorf("error capturing stderr: %v", err)
		}
		if !bytes.Contains([]byte(stdErr), []byte(testCase.message)) {
			t.Errorf("expected stderr to contain %q, but got %q", testCase.message, stdErr)
		}
		if !bytes.Contains([]byte(stdErr), []byte(testCase.expectedLogLevel)) {
			t.Errorf("expected stderr to contain %q, but got %q", testCase.expectedLogLevel, stdErr)
		}
	}
}

func TestLogln(t *testing.T) {
	testCases := []struct {
		message          string
		logLevel         glog.LogLevel
		expectedLogLevel string
	}{
		{
			message:          "test infoln",
			logLevel:         glog.LogLevelInfo,
			expectedLogLevel: "I0224",
		},
		{
			message:          "test warnln",
			logLevel:         glog.LogLevelWarn,
			expectedLogLevel: "W0224",
		},
		{
			message:          "test errorln",
			logLevel:         glog.LogLevelError,
			expectedLogLevel: "E0224",
		},
	}

	for _, testCase := range testCases {
		stdErr, err := captureStderr(func() {
			glog.Logln(testCase.logLevel, testCase.message)
		})
		if err != nil {
			t.Errorf("error capturing stderr: %v", err)
		}
		if !bytes.Contains([]byte(stdErr), []byte(testCase.message)) {
			t.Errorf("expected stderr to contain %q, but got %q", testCase.message, stdErr)
		}
		if !bytes.Contains([]byte(stdErr), []byte(testCase.expectedLogLevel)) {
			t.Errorf("expected stderr to contain %q, but got %q", testCase.expectedLogLevel, stdErr)
		}
	}
}
