package logger

import (
	"testing"
)

func Test_Logger(t *testing.T) {
	Debug("0000000000")

	logger := NewLogger(WithLevel("info"))
	SetLogger(logger)
	Debug("1111111111")
}
