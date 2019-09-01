/*
 * Copyright (c) 2019. 陈金龙.
 */

package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FileLogger struct {
	zapLogger *zap.Logger
}
func NewFileLogger(filePath string, maxSize int, maxBackups int, maxAge int) *FileLogger {
	hook := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, //days
		Compress:   false,  // disabled by default
	}

	levelFunc := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.InfoLevel
	})

	fileWriter := zapcore.AddSync(&hook)
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, levelFunc)

	core := zapcore.NewTee(
		fileCore,
	)

	return &FileLogger{
		zap.New(core),
	}
}

func (fileLogger *FileLogger)Panicf(fmt string, args ...interface{}) {
	write(fileLogger.zapLogger, LevelPanic, fmt, args...)
}
func (fileLogger *FileLogger)Fatalf(fmt string, args ...interface{}) {
	write(fileLogger.zapLogger, LevelFatal, fmt, args...)
}
func (fileLogger *FileLogger)Errorf(fmt string, args ...interface{}) {
	write(fileLogger.zapLogger, LevelError, fmt, args...)
}
func (fileLogger *FileLogger)Infof(fmt string, args ...interface{}) {
	write(fileLogger.zapLogger, LevelInfo, fmt, args...)
}
func (fileLogger *FileLogger)Debugf(fmt string, args ...interface{}) {
	write(fileLogger.zapLogger, LevelDebug, fmt, args...)
}
func (fileLogger *FileLogger)Warnf(fmt string, args ...interface{}) {
	write(fileLogger.zapLogger, LevelWarn, fmt, args...)
}
func (fileLogger *FileLogger)Messagef(level Level, fmt string, args ...interface{}) {
	write(fileLogger.zapLogger, level, fmt, args...)
}
func (fileLogger *FileLogger)Flush() {
	if fileLogger.zapLogger != nil {
		fileLogger.zapLogger.Sync()
	}
}
