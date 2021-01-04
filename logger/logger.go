package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

type Fields map[string]interface{}

var log *logrus.Logger

var f *logrus.TextFormatter

func init() {
	log = logrus.New()
	f = &logrus.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
		FullTimestamp:    true,
		DisableColors:    false,
	}
	log.SetFormatter(f)
	dir, _ := os.Getwd()
	errFile, err := os.OpenFile(dir+"/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open log file failed：", err)
	}
	hook := ErrorHook{Writer: errFile}
	log.AddHook(&hook)
}

func Info(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Print(getFileInfo() + getLogMsg(msg, "37"))
}

func FInfo(format string, args ...interface{}) {
	log.Info(getFileInfo() + fmt.Sprintf(format, args...))
}

func Success(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Print(getFileInfo() + getLogMsg(msg, "32"))
}

func FSuccess(msg string, args ...interface{}) {
	log.Print(getFileInfo() + getLogMsg(fmt.Sprintf(msg, args...), "32"))
}

func Warning(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Warning(getFileInfo() + getLogMsg(msg, "33"))
}

func FWarning(msg string, args ...interface{}) {
	log.Warning(getFileInfo() + getLogMsg(fmt.Sprintf(msg, args...), "33"))
}

func Error(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Error(getFileInfo() + getLogMsg(msg, "31"))
}
func FError(msg string, args ...interface{}) {
	log.Error(getFileInfo() + getLogMsg(fmt.Sprintf(msg, args...), "31"))
}

type CronLogger struct {
}

func (l *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	v := []interface{}{"cron info", msg, keysAndValues}
	log.Println(getLogContent("", v))
}

func (l *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	v := []interface{}{"cron error", err, msg, keysAndValues}
	log.Println(getLogContent("31", v))
}

func getLogMsg(msg string, color string) string {
	return "\x1b[0;" + color + "m" + msg + "\x1b[0m"
}

func getLogContent(color string, v []interface{}) []interface{} {
	var content []interface{}
	if color != "" {
		c := "\x1b[0;" + color + "m"
		content = append(content, c)
	}
	content = append(content, getFileInfo())
	for _, value := range v {
		content = append(content, value)
	}
	if color != "" {
		content = append(content, "\x1b[0m")
	}
	return content
}

func getFileInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return fmt.Sprintf("%s:%d:", file, line)
}
