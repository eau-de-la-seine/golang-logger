package logger

import (
	"reflect"
	"io"
	"os"
	"time"
	"runtime"
	"fmt"
)

const (
	ERROR = iota
	WARN
	INFO
	DEBUG
)

var stringLevels []string = []string{"ERROR", "WARN", "INFO", "DEBUG"}

type Logger struct {
	yourType reflect.Type
	writer io.Writer
	level int
}

func (logger *Logger) log(file string, line int, level int, messageFormat string, v ...interface{}) {
	if !(logger.level <= level) {
		return
	}

	formattedMessage := fmt.Sprintf(messageFormat, v...)
	fmt.Fprintf(logger.writer, "[%s][%s][" + logger.yourType.Name() + "][%s:%d] : %s\n", time.Now().Format(time.RFC3339), stringLevels[logger.level], file, line, formattedMessage)
}

func (logger *Logger) Debug(messageFormat string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	logger.log(file, line, DEBUG, messageFormat, v...)
}

func (logger *Logger) Info(messageFormat string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	logger.log(file, line, INFO, messageFormat, v...)
}

func (logger *Logger) Warn(messageFormat string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	logger.log(file, line, WARN, messageFormat, v...)
}

func (logger *Logger) Error(messageFormat string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	logger.log(file, line, ERROR, messageFormat, v...)
}

func NewLogger(yourType reflect.Type, writer io.Writer, level int) (*Logger) {
	if yourType == nil {
		panic("[NewLogger] Parameter `yourType` must not be nil")
	}

	if writer == nil {
		panic("[NewLogger] Parameter `writer` must not be nil")
	}

	if !(level >= ERROR && level <= DEBUG) {
		panic("[NewLogger] Parameter `level` must be ERROR, WARN, INFO, or DEBUG")
	}

	var logger *Logger = new(Logger)
	logger.yourType = yourType
	logger.writer = writer
	logger.level = level
	return logger
}

func NewConsoleLogger(yourType reflect.Type, level int) (*Logger) {
	return NewLogger(yourType, os.Stdout, level)
}
