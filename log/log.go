package log

import (
	"encoding/json"
	"fmt"
	"github.com/jinlongchen/golang-utilities/converter"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
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
	Config("golang-util",
		LevelDebug,
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
		consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     func(t time.Time, enc zapcore.PrimitiveArrayEncoder){
				enc.AppendString(t.Local().Format("2006-01-02 15:04:05"))
			},
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
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
		fileEncoder := zapcore.NewJSONEncoder(newProductionEncoderConfig())
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
	write(globalZapLogger, LevelPanic, fmt, args...)
}
func Fatalf(fmt string, args ...interface{}) {
	write(globalZapLogger, LevelFatal, fmt, args...)
}
func Errorf(fmt string, args ...interface{}) {
	write(globalZapLogger, LevelError, fmt, args...)
}
func Infof(fmt string, args ...interface{}) {
	write(globalZapLogger, LevelInfo, fmt, args...)
}
func Debugf(fmt string, args ...interface{}) {
	write(globalZapLogger, LevelDebug, fmt, args...)
}
func Warnf(fmt string, args ...interface{}) {
	write(globalZapLogger, LevelWarn, fmt, args...)
}
func Messagef(level Level, fmt string, args ...interface{}) {
	write(globalZapLogger, level, fmt, args...)
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
	write(globalZapLogger, LevelInfo, "%s", string(jd))
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
func write(zapLogger *zap.Logger, level Level, format string, args ...interface{}) {
	if zapLogger == nil {
		zapLogger = globalZapLogger
	}
	if zapLogger == nil {
		return
	}

	fields := make([]zap.Field, 0)
	if globalAppName != "" {
		fields = append(fields, zap.String("app", globalAppName))
	}
	if level == LevelPanic {
		zapLogger.Panic(fmt.Sprintf(format, args...), fields...)
		os.Exit(1)
	} else if level == LevelFatal {
		zapLogger.Fatal(fmt.Sprintf(format, args...), fields...)
		os.Exit(1)
	} else if level == LevelError {
		zapLogger.Error(fmt.Sprintf(format, args...), fields...)
	} else if level == LevelWarn {
		zapLogger.Warn(fmt.Sprintf(format, args...), fields...)
	} else if level == LevelInfo {
		zapLogger.Info(fmt.Sprintf(format, args...), fields...)
	} else if level == LevelDebug {
		zapLogger.Debug(fmt.Sprintf(format, args...), fields...)
	}
}

func newProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
