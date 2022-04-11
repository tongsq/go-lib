package logger

import (
	"encoding/json"
	"io"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tongsq/go-lib/util"
)

type ErrorHook struct {
	Writer io.Writer
}

func (hook *ErrorHook) Fire(entry *logrus.Entry) error {

	var data logrus.Fields = util.MapCopy(entry.Data)
	msg := entry.Message
	msg = strings.TrimRight(msg, "\u001B[0m")
	msg = strings.TrimLeft(msg, "\u001b[0;")
	reg := regexp.MustCompile(`^[0-9]+m`)
	msg = reg.ReplaceAllString(msg, "")
	data["message"] = msg
	data["time"] = entry.Time
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
