package log

import (
	"fmt"
	"os"
)

type logger struct {
	level Level
}

var log logger

func InitLogger(level Level) {
	log = logger{
		level: level,
	}
	Debug("Set log level to:", LevelToString[level])
}

func Debug(args ...interface{}) {
	Print(LevelDebug, args...)
}

func Info(args ...interface{}) {
	Print(LevelInfo, args...)
}

func Warn(args ...interface{}) {
	Print(LevelWarn, args...)
}

func Fatal(args ...interface{}) {
	Print(LevelFatal, args...)
}

func Print(logLevel Level, args ...interface{}) {
	print(logLevel, append([]interface{}{"[", LevelToString[logLevel], "]"}, args...)...)
}

func Stream(logLevel Level, args ...interface{}) {
	print(logLevel, append([]interface{}{"          "}, args...)...)
}

func print(logLevel Level, args ...interface{}) {
	if log.level <= logLevel {
		fmt.Println(args...)
	}

	if logLevel == LevelFatal {
		os.Exit(1)
	}
}
