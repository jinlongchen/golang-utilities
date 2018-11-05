package log

import (
	"github.com/sirupsen/logrus"
	"encoding/json"
)

type LogLevel string
type LogFormat string

var (
	appName string
	logger = logrus.New()
)

const (
	LOG_FORMAT_JSON LogFormat = "json"
	LOG_FORMAT_TEXT LogFormat = "text"

	LOG_LEVEL_DEBUG LogLevel = "debug"
	LOG_LEVEL_INFO LogLevel = "info"
	LOG_LEVEL_WARN LogLevel = "warn"
	LOG_LEVEL_ERROR LogLevel = "error"
	LOG_LEVEL_FATAL LogLevel = "fatal"
)

func InitLogger(app string, level LogLevel, format LogFormat, disableColors bool) *logrus.Logger {
	appName = app

	lvl := logrus.InfoLevel

	switch level {
	case LOG_LEVEL_DEBUG:
		lvl = logrus.DebugLevel
	case LOG_LEVEL_INFO:
		lvl = logrus.InfoLevel
	case LOG_LEVEL_WARN:
		lvl = logrus.WarnLevel
	case LOG_LEVEL_ERROR:
		lvl = logrus.ErrorLevel
	case LOG_LEVEL_FATAL:
		lvl = logrus.FatalLevel
	default:
		lvl = logrus.InfoLevel
	}

	logger.SetLevel(lvl)

	if format == "json" {
		logger.Formatter = new(logrus.JSONFormatter)
	} else {
		logger.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: disableColors}
	}
	return logger
}

func AddHook(hook logrus.Hook) {
	logger.AddHook(hook)
}
func DefaultLogger() *logrus.Logger {
	return logger
}
func Panicf(fmt string, args ...interface{}) {
	write(logrus.PanicLevel, fmt, args...)
}
func Fatalf(fmt string, args ...interface{}) {
	write(logrus.FatalLevel, fmt, args...)
}
func Errorf(fmt string, args ...interface{}) {
	write(logrus.ErrorLevel, fmt, args...)
}
func Infof(fmt string, args ...interface{}) {
	write(logrus.InfoLevel, fmt, args...)
}
func Debugf(fmt string, args ...interface{}) {
	write(logrus.DebugLevel, fmt, args...)
}
func Warnf(fmt string, args ...interface{}) {
	write(logrus.WarnLevel, fmt, args...)
}
func Json(j interface{}) {
	jd, err := json.Marshal(j)
	if err != nil {
		return
	}
	write(logrus.InfoLevel, "%s", string(jd))
}

func DumpFormattedJson(j interface{}) {
	jd, err := json.MarshalIndent(j,"", " ")
	if err != nil {
		return
	}
	println(string(jd))
}

func write(level logrus.Level, fmt string, args ...interface{}) {
	callstack := string(stack(0))
	if level == logrus.PanicLevel {
		logger.WithField("app", appName).WithField("caller", callstack).Panicf(fmt, args...)
	} else if level == logrus.FatalLevel {
		logger.WithField("app", appName).WithField("caller", callstack).Fatalf(fmt, args...)
	} else if level == logrus.ErrorLevel {
		logger.WithField("app", appName).WithField("caller", callstack).Errorf(fmt, args...)
	} else if level == logrus.WarnLevel {
		logger.WithField("app", appName).WithField("caller", callstack).Warnf(fmt, args...)
	} else if level == logrus.InfoLevel {
		logger.WithField("app", appName).Infof(fmt, args...)
	} else if level == logrus.DebugLevel {
		logger.WithField("app", appName).Debugf(fmt, args...)
	}
}
