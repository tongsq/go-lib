package tests

import (
	"github.com/tongsq/go-lib/logger"
	"testing"
)

func TestLog(t *testing.T) {
	logger.FInfo("test logger FInfo, %s", "hello world")
	logger.FWarning("test logger FWarning, %s", "hello world")
	logger.FError("test logger FError, %s", "hello world")
	logger.FSuccess("test logger FSuccess, %s", "hello world")
	logger.Info("test logger Info", logger.Fields{"key": "test"})
	logger.Warning("test logger Warning", logger.Fields{"key": "test"})
	logger.Error("test logger Error", logger.Fields{"key": "test"})
	logger.Success("test logger Success", logger.Fields{"key": "test"})
	t.Log("TestAdd success")
}
