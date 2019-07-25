package log

import (
	"encoding/json"
	"fmt"
	"github.com/jinlongchen/golang-utilities/converter"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Level string
type Format string

var (
	globalZapLogger *zap.Logger
	//globalLogrusLogger = logrus.New()
	globalAppName string
	globalLevel   Level
	globalFormat  Format
)

const (
	FormatJSON Format = "json"
	FormatText Format = "text"

	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelFatal Level = "fatal"
	LevelPanic Level = "panic"
)

func init() {
	Config("update location",
		Level("debug"),
		true,
		"",
		0,
		0,
		0,
	)
}

func Config(appName string, level Level, console bool, filename string, maxSize int, maxBackups int, maxAge int) {
	globalAppName = appName
	globalLevel = level

	levels := map[Level]zapcore.Level{
		LevelDebug: zap.DebugLevel,
		LevelInfo:  zap.InfoLevel,
		LevelWarn:  zap.WarnLevel,
		LevelError: zap.ErrorLevel,
		LevelFatal: zap.PanicLevel,
		LevelPanic: zap.FatalLevel,
	}

	zaplvl := levels[level]

	levelFunc := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zaplvl
	})

	cores := make([]zapcore.Core, 0)

	if console {
		consoleDebugging := zapcore.Lock(os.Stdout)
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleCore := zapcore.NewCore(consoleEncoder, consoleDebugging, levelFunc)

		cores = append(cores, consoleCore)
	}
	if filename != "" {
		hook := lumberjack.Logger{
			Filename:   filename,
			MaxSize:    maxSize, // megabytes
			MaxBackups: maxBackups,
			MaxAge:     maxAge, //days
			Compress:   false,  // disabled by default
		}

		fileWriter := zapcore.AddSync(&hook)
		fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		fileCore := zapcore.NewCore(fileEncoder, fileWriter, levelFunc)

		cores = append(cores, fileCore)
	}

	core := zapcore.NewTee(
		cores...,
	)

	logger := zap.New(core, zap.AddStacktrace(zap.ErrorLevel))

	if globalZapLogger != nil {
		globalZapLogger.Sync()
	}
	globalZapLogger = logger
}

//func Config(appName string, level Level, format Format) {
//	globalAppName = appName
//	globalLevel = level
//	globalFormat = format
//
//	switch level {
//	case LevelNone:
//		globalLogrusLogger = nil
//	case LevelDebug:
//		globalLogrusLogger.SetLevel(logrus.DebugLevel)
//	case LevelInfo:
//		globalLogrusLogger.SetLevel(logrus.InfoLevel)
//	case LevelWarn:
//		globalLogrusLogger.SetLevel(logrus.WarnLevel)
//	case LevelError:
//		globalLogrusLogger.SetLevel(logrus.ErrorLevel)
//	case LevelFatal:
//		globalLogrusLogger.SetLevel(logrus.FatalLevel)
//	case LevelPanic:
//		globalLogrusLogger.SetLevel(logrus.PanicLevel)
//	default:
//		globalLogrusLogger.SetLevel(logrus.InfoLevel)
//	}
//
//	if format == FormatJSON {
//		if globalLogrusLogger != nil {
//			globalLogrusLogger.Formatter = new(logrus.JSONFormatter)
//		}
//	} else {
//		if globalLogrusLogger != nil {
//			globalLogrusLogger.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: true}
//		}
//	}
//}

//func SetOutput(out io.Writer) {
//	if out == nil {
//		return
//	}
//	globalLogrusLogger.SetOutput(out)
//}
//func AddHook(hook logrus.Hook) {
//	globalLogrusLogger.AddHook(hook)
//}
func Panicf(fmt string, args ...interface{}) {
	write(LevelPanic, fmt, args...)
}
func Fatalf(fmt string, args ...interface{}) {
	write(LevelFatal, fmt, args...)
}
func Errorf(fmt string, args ...interface{}) {
	write(LevelError, fmt, args...)
}
func Infof(fmt string, args ...interface{}) {
	write(LevelInfo, fmt, args...)
}
func Debugf(fmt string, args ...interface{}) {
	write(LevelDebug, fmt, args...)
}
func Warnf(fmt string, args ...interface{}) {
	write(LevelWarn, fmt, args...)
}
func Messagef(level Level, fmt string, args ...interface{}) {
	write(level, fmt, args...)
}
func Flush() {
	if globalZapLogger != nil {
		globalZapLogger.Sync()
	}
}
func Json(j interface{}) {
	jd, err := json.Marshal(j)
	if err != nil {
		return
	}
	write(LevelInfo, "%s", string(jd))
}
func DumpFormattedJson(j interface{}) {
	jd, err := json.MarshalIndent(j, "", " ")
	if err != nil {
		return
	}
	println(string(jd))
}
func DumpKeyValue(j interface{}) {
	m := converter.ConvertToMap(j)
	for key, value := range m {
		fmt.Printf("%s:%v\n", key, value)
	}
}
func write(level Level, format string, args ...interface{}) {
	//if globalLogrusLogger == nil {
	//	return
	//}
	////callStack := string(stack(0))
	//if level == LevelPanic {
	//	globalLogrusLogger.WithField("app", globalAppName).Panicf(format, args...)
	//} else if level == LevelFatal {
	//	globalLogrusLogger.WithField("app", globalAppName).Fatalf(format, args...)
	//} else if level == LevelError {
	//	globalLogrusLogger.WithField("app", globalAppName).Errorf(format, args...)
	//} else if level == LevelWarn {
	//	globalLogrusLogger.WithField("app", globalAppName).Warnf(format, args...)
	//} else if level == LevelInfo {
	//	globalLogrusLogger.WithField("app", globalAppName).Infof(format, args...)
	//} else if level == LevelDebug {
	//	globalLogrusLogger.WithField("app", globalAppName).Debugf(format, args...)
	//}
	if globalZapLogger == nil {
		return
	}

	if level == LevelPanic {
		globalZapLogger.Panic(fmt.Sprintf(format, args...), zap.String("app", globalAppName))
		os.Exit(1)
	} else if level == LevelFatal {
		globalZapLogger.Fatal(fmt.Sprintf(format, args...), zap.String("app", globalAppName))
		os.Exit(1)
	} else if level == LevelError {
		globalZapLogger.Error(fmt.Sprintf(format, args...), zap.String("app", globalAppName))
	} else if level == LevelWarn {
		globalZapLogger.Warn(fmt.Sprintf(format, args...), zap.String("app", globalAppName))
	} else if level == LevelInfo {
		globalZapLogger.Info(fmt.Sprintf(format, args...), zap.String("app", globalAppName))
	} else if level == LevelDebug {
		globalZapLogger.Debug(fmt.Sprintf(format, args...), zap.String("app", globalAppName))
	}
}
