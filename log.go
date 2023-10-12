package main

import (
	"fmt"
	"os"
)

type Logger struct {
	IsDebug bool
}

var log = &Logger{
	IsDebug: false,
}

func (t *Logger) Debug(v ...interface{}) {
	fmt.Println(v...)
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
