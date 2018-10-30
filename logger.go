package logger

import (
	"reflect"
	"io"
	"os"
	"time"
	"runtime"
	"fmt"
	"path"
)

const (
	LEVEL_ERROR = iota
	LEVEL_WARN
	LEVEL_INFO
	LEVEL_DEBUG
)

var stringLevels []string = []string{"ERROR", "WARN", "INFO", "DEBUG"}

type Logger struct {
	yourType reflect.Type
	writer io.Writer
	level int
}

func (logger *Logger) log(level int, messageFormat string, v ...interface{}) {
	if logger.level < level {
		return
	}

	_, filePath, lineNumber, _ := runtime.Caller(2)
	formattedMessage := fmt.Sprintf(messageFormat, v...)
	fmt.Fprintf(logger.writer,
		"[%s][%s][" + logger.yourType.Name() + "][%s:%d] : %s\n",
		time.Now().Format(time.RFC3339),
		stringLevels[level],
		path.Base(filePath),
		lineNumber,
		formattedMessage)
}

func (logger *Logger) Debug(messageFormat string, v ...interface{}) {
	logger.log(LEVEL_DEBUG, messageFormat, v...)
}

func (logger *Logger) Info(messageFormat string, v ...interface{}) {
	logger.log(LEVEL_INFO, messageFormat, v...)
}

func (logger *Logger) Warn(messageFormat string, v ...interface{}) {
	logger.log(LEVEL_WARN, messageFormat, v...)
}

func (logger *Logger) Error(messageFormat string, v ...interface{}) {
	logger.log(LEVEL_ERROR, messageFormat, v...)
}

func NewLogger(yourType reflect.Type, writer io.Writer, level int) (*Logger) {
	if yourType == nil {
		panic("[NewLogger] Parameter `yourType` must not be nil")
	}

	if writer == nil {
		panic("[NewLogger] Parameter `writer` must not be nil")
	}

	if !(level >= LEVEL_ERROR && level <= LEVEL_DEBUG) {
		panic("[NewLogger] Parameter `level` must be LEVEL_ERROR, LEVEL_WARN, LEVEL_INFO, or LEVEL_DEBUG")
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
