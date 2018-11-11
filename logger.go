package logger

import (
	"io"
	"os"
	"time"
	"runtime"
	"fmt"
	"path"
	"sync"
)

const (
	LEVEL_ERROR = iota
	LEVEL_WARN
	LEVEL_INFO
	LEVEL_DEBUG
)

var stringLevels []string = []string{"ERROR", "WARN", "INFO", "DEBUG"}

type Logger struct {
	writer io.Writer
	level int
	mutex sync.Mutex
}

func (logger *Logger) log(level int, messageFormat string, v ...interface{}) {
	if logger.level < level {
		return
	}

	_, filePath, lineNumber, _ := runtime.Caller(2)

	formattedLogMessage := fmt.Sprintf("[%s][%s][%s:%d] : %s\n",
		time.Now().Format(time.RFC3339),
		stringLevels[level],
		// Only show the file name from given path
		path.Base(filePath),
		lineNumber,
		// Formatted user message
		fmt.Sprintf(messageFormat, v...))

	messageToLog := []byte(formattedLogMessage)
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	logger.writer.Write(messageToLog)
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

func NewLogger(writer io.Writer, level int) (*Logger) {
	if writer == nil {
		panic("[NewLogger] Parameter `writer` must not be nil")
	}

	if !(level >= LEVEL_ERROR && level <= LEVEL_DEBUG) {
		panic("[NewLogger] Parameter `level` must be LEVEL_ERROR, LEVEL_WARN, LEVEL_INFO, or LEVEL_DEBUG")
	}

	var logger *Logger = new(Logger)
	logger.writer = writer
	logger.level = level
	return logger
}

func NewConsoleLogger(level int) (*Logger) {
	return NewLogger(os.Stdout, level)
}
