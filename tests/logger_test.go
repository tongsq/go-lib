package tests

import (
	"github.com/tongsq/go-lib/logger"
	"testing"
)

func TestLog(t *testing.T) {
	logger.Info("test logger info", "hello world")
	logger.Warning("test logger warning", "hello world")
	logger.Error("test logger error", "hello world")
	logger.Success("test logger success", "hello world")
	t.Log("TestAdd success")
}
