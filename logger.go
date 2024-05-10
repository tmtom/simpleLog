package simpleLog

// Simple global logger supporting levels. Uses standard log underneath

import (
	"fmt"
	"io"
	gLog "log"
	"os"
)

type LogLevel int

const (
	ERROR LogLevel = 5 * iota
	WARN
	INFO
	DEBUG
)

var (
	logger   *gLog.Logger
	level    = INFO
	useColor = true
)

var MsgLevel = map[LogLevel]string{
	ERROR: "ERRO",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBG",
}

var MsgLevelC = map[LogLevel]string{
	ERROR: "\033[1;31mERRO\033[0m",
	WARN:  "\033[1;33mWARN\033[0m",
	INFO:  "\033[1;34mINFO\033[0m",
	DEBUG: "\033[0;36mDEBG\033[0m",
}

// Sets new output (also recreates underlying logger)
func SetOutput(out io.Writer) {
	logger = gLog.New(out, "", gLog.Lmicroseconds|gLog.Lshortfile)
}

func SetLevel(lvl LogLevel) {
	level = lvl
}

func init() {
	SetOutput(os.Stderr)
	SetLevel(INFO)
}

func logInternal(logLevel LogLevel, msg string) {
	if level >= logLevel {
		var m *map[LogLevel]string
		if useColor {
			m = &MsgLevelC
		} else {
			m = &MsgLevel
		}
		logger.Output(3, (*m)[logLevel]+": "+msg)
	}
}

func logInternalF(logLevel LogLevel, format string, v ...interface{}) {
	if level >= logLevel {
		var m *map[LogLevel]string
		if useColor {
			m = &MsgLevelC
		} else {
			m = &MsgLevel
		}
		logger.Output(3, fmt.Sprintf((*m)[logLevel]+": "+format, v...))
	}
}

func Error(msg string) {
	logInternal(ERROR, msg)
}

func Errorf(format string, v ...interface{}) {
	logInternalF(ERROR, format, v...)
}

func Warn(msg string) {
	logInternal(WARN, msg)
}

func Warnf(format string, v ...interface{}) {
	logInternalF(WARN, format, v...)
}

func Info(msg string) {
	logInternal(INFO, msg)
}

func Infof(format string, v ...interface{}) {
	logInternalF(INFO, format, v...)
}

func Debug(msg string) {
	logInternal(DEBUG, msg)
}

func Debugf(format string, v ...interface{}) {
	logInternalF(DEBUG, format, v...)
}
