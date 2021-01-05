package main

import (
	"github.com/tongsq/go-lib/logger"
)

func main() {
	//logrus.SetFormatter(&logrus.TextFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//	DisableTimestamp: false,
	//	FullTimestamp: true,
	//})
	//logrus.WithFields(logrus.Fields{"key":"value"}).Info("hello world")
	logger.Info("hello world", logger.Fields{"hello": "world", "key": "aaa"})
	logger.Warning("hello world", logger.Fields{"hello": "world", "key": "aaa"})
	logger.Error("hello world", logger.Fields{"hello": "world", "key": "aaa"})
	logger.Success("hello world", logger.Fields{"hello": "world", "key": "aaa"})
	logger.FInfo("aa", "bb")
	logger.FError("aa, %s", "bb")
}
