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
	LEVEL_OFF
)

var stringLevels []string = []string{"ERROR", "WARN", "INFO", "DEBUG"}

type Logger struct {
	writer io.Writer
	level int
	mutex sync.Mutex
}

func (logger *Logger) log(level int, messageFormat string, v ...interface{}) {
	if logger.level == LEVEL_OFF || logger.level < level {
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

	if !(level >= LEVEL_ERROR && level <= LEVEL_OFF) {
		panic("[NewLogger] Parameter `level` must be LEVEL_ERROR, LEVEL_WARN, LEVEL_INFO, LEVEL_DEBUG, or LEVEL_OFF")
	}

	var logger *Logger = new(Logger)
	logger.writer = writer
	logger.level = level
	return logger
}

func NewConsoleLogger(level int) (*Logger) {
	return NewLogger(os.Stdout, level)
}

// Logger factory (experimental)
type LoggerOptions struct {
	LoggerLevel int
}

var globalLoggerOptions LoggerOptions = LoggerOptions{
	LoggerLevel: LEVEL_OFF}

func SetLoggerFactoryOptions(loggerOptions LoggerOptions) {
	globalLoggerOptions = loggerOptions
}

var loggerMap sync.Map

func LoggerFactory(name string, writer io.Writer) (*Logger) {
	actualLogger, _ := loggerMap.LoadOrStore(name, NewLogger(writer, globalLoggerOptions.LoggerLevel))
	return actualLogger.(*Logger)
}
