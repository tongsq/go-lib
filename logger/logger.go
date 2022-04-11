package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	green  = 32
	yellow = 33
	blue   = 36
	gray   = 37
)

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

// Level type
type Level uint32

type Fields map[string]interface{}

var log *logrus.Logger

var f *logrus.TextFormatter

var errHook *ErrorHook

func init() {
	log = logrus.New()
	f = &logrus.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		DisableTimestamp:          false,
		FullTimestamp:             true,
		ForceColors:               true,
		DisableColors:             false,
		EnvironmentOverrideColors: false,
		PadLevelText:              false,
	}
	log.SetFormatter(f)
}

func SetErrorFile(path string) {
	errFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open log file failed：", err)
	}
	if errHook == nil {
		errHook = &ErrorHook{Writer: errFile}
		log.AddHook(errHook)
	} else {
		errHook.Writer = errFile
	}
}

func SetLogLevel(l Level) {
	switch l {
	case TraceLevel:
		log.SetLevel(logrus.TraceLevel)
	case DebugLevel:
		log.SetLevel(logrus.DebugLevel)
	case InfoLevel:
		log.SetLevel(logrus.InfoLevel)
	case WarnLevel:
		log.SetLevel(logrus.WarnLevel)
	case ErrorLevel:
		log.SetLevel(logrus.ErrorLevel)
	case FatalLevel:
		log.SetLevel(logrus.FatalLevel)
	case PanicLevel:
		log.SetLevel(logrus.PanicLevel)
	default:
		log.Error("SetLogLevel fail , unknown  level")
	}
}

func GetLogger() *logrus.Logger {
	return log
}

func Debug(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Debug(getFileInfo() + msg)
}

func FDebug(msg string, args ...interface{}) {
	log.Debug(getFileInfo() + fmt.Sprintf(msg, args...))
}

func Info(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Print(getFileInfo() + getLogMsg(msg, gray))
}

func FInfo(format string, args ...interface{}) {
	log.Info(getFileInfo() + fmt.Sprintf(format, args...))
}

func Success(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Print(getFileInfo() + getLogMsg(msg, blue))
}

func FSuccess(msg string, args ...interface{}) {
	log.Print(getFileInfo() + getLogMsg(fmt.Sprintf(msg, args...), blue))
}

func Warning(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Warning(getFileInfo() + getLogMsg(msg, yellow))
}

func FWarning(msg string, args ...interface{}) {
	log.Warning(getFileInfo() + getLogMsg(fmt.Sprintf(msg, args...), yellow))
}

func Error(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Error(getLogMsg(getFileInfo()+msg, red))
}
func FError(msg string, args ...interface{}) {
	log.Error(getLogMsg(getFileInfo()+fmt.Sprintf(msg, args...), red))
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

func getLogMsg(msg string, color int) string {
	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, msg)
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
