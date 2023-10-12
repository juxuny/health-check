package log

import (
	"fmt"
	"os"
)

type Logger struct {
	IsDebug bool
}

var defaultInstance = &Logger{
	IsDebug: false,
}

func (t *Logger) Debug(v ...interface{}) {
	if t.IsDebug {
		fmt.Println(v...)
	}
}

func (t *Logger) Info(v ...interface{}) {
	fmt.Println(v...)
}

func (t *Logger) Error(v ...interface{}) {
	fmt.Println(v...)
}

func (t *Logger) Fatal(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(255)
}

func Debug(v ...interface{}) {
	defaultInstance.Debug(v...)
}

func Info(v ...interface{}) {
	defaultInstance.Info(v...)
}

func Error(v ...interface{}) {
	defaultInstance.Error(v...)
}

func Fatal(v ...interface{}) {
	defaultInstance.Fatal(v...)
}

func SetDebug(verbose bool) {
	defaultInstance.IsDebug = verbose
}
