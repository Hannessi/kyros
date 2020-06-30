package log

import (
	"fmt"
	"time"
)

const logColorInfo = "\033[0;37m%s\033[0m"
const logColorWarn = "\033[0;33m%s\033[0m"
const logColorError = "\033[0;31m%s\033[0m"
const logColorDebug = "\033[0;35m%s\033[0m"

const logLevelInfo = "Info"
const logLevelDebug = "Debug"
const logLevelError = "Error"
const logLevelWarn = "Warning"

func Info(messages ...interface{}) {
	printToStdout(messages, logLevelInfo, logColorInfo)
}

func Debug(messages ...interface{}) {
	printToStdout(messages, logLevelDebug, logColorDebug)
}

func Error(messages ...interface{}) {
	printToStdout(messages, logLevelError, logColorError)
}

func Warn(messages ...interface{}) {
	printToStdout(messages, logLevelWarn, logColorWarn)
}

func Fatal(messages ...interface{}) {
	printToStdout(messages, logLevelError, logColorError)
}

func printToStdout(message interface{}, level, color string) {
	now := time.Now()
	_, _ = fmt.Printf(color, fmt.Sprintf("time: '%s' level: '%s' message:'%s'\n", now.String(), level, message))
}
