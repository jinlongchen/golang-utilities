/*
 * Copyright (c) 2019. 陈金龙.
 */

package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewFileLogger(filePath string,
	maxSize int,
	maxBackups int,
	maxAge int) *zap.Logger {

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

	return zap.New(core)
}
