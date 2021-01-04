package logger

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/tongsq/go-lib/util"
	"io"
)

type ErrorHook struct {
	Writer io.Writer
}

func (hook *ErrorHook) Fire(entry *logrus.Entry) error {

	var data logrus.Fields = util.MapCopy(entry.Data)
	data["Message"] = entry.Message
	data["Time"] = entry.Time
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	jsonStr = append(jsonStr, '\n')
	hook.Writer.Write(jsonStr)
	return nil
}

func (hook *ErrorHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}
