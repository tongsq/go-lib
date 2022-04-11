package tests

import (
	"testing"

	"github.com/tongsq/go-lib/logger"
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

func TestSetLevel(t *testing.T) {
	logger.SetLogLevel(4)
	logger.Info("this is info log", map[string]interface{}{})
	logger.Warning("this is warn log", map[string]interface{}{})
}

func TestLogError(t *testing.T) {
	logger.SetErrorFile("err.log")
	logger.Error("this error log", map[string]interface{}{"data": "abc"})
}
