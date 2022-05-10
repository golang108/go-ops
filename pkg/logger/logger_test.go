package logger

import "testing"

func TestLogger(t *testing.T) {
	Info("info msg")
	Debug("debug msg")
	Warn("warn msg")
	Error("error msg")
}

func TestLoggerInit(t *testing.T) {
	InitLog(nil)
	Info("info msg")
	Debug("debug msg")
	Warn("warn msg")
	Error("error msg")
}
