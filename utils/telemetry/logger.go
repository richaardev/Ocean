package telemetry

import (
	"fmt"
	"time"
)

func log(message string, level string) {
	tm := time.Now().UTC().Format("2006-01-02 15:04:05")

	fmt.Printf("\x1b[37m[%s] \x1b[0m%s\x1b[0m: %s\n", tm, level,message)
}

func Info(message string) {
	go log(message, "\x1b[34;1mINFO")
}
func Infof(message string, args ...interface{}) {
	go Info(fmt.Sprintf(message, args...))
}

func Warn(message string) {
	go log(message, "\x1b[33;1mWARN")
}
func Warnf(message string, args ...interface{}) {
	go Warn(fmt.Sprintf(message, args...))
}

func Error(message string) {
	go log(message, "\x1b[31;1mERROR")
}
func Errorf(message string, args ...interface{}) {
	go Error(fmt.Sprintf(message, args...))
}

func Debug(message string) {
	go log(message, "\x1b[35;1mDEBUG")
}
func Debugf(message string, args ...interface{}) {
	go Debug(fmt.Sprintf(message, args...))
}