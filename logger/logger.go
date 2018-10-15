package logger

import (
	"reflect"
	"io"
	"os"
	"time"
	"runtime"
	"fmt"
)

type Logger struct {
	yourType reflect.Type
	debugWriter, infoWriter, warnWriter, errorWriter io.Writer
}

func (logger *Logger) Debug(messageFormat string, v ...interface{}) {
	formattedMessage := fmt.Sprintf(messageFormat, v...)
	_, file, line, _ := runtime.Caller(1)
	fmt.Fprintf(logger.debugWriter, "[%s][DEBUG][" + logger.yourType.Name() + "][%s:%d] : %s\n", time.Now().Format(time.RFC3339), file, line, formattedMessage)
}

func (logger *Logger) Info(messageFormat string, v ...interface{}) {
	formattedMessage := fmt.Sprintf(messageFormat, v...)
	_, file, line, _ := runtime.Caller(1)
	fmt.Fprintf(logger.infoWriter, "[%s][INFO][" + logger.yourType.Name() + "][%s:%d] : %s\n", time.Now().Format(time.RFC3339), file, line, formattedMessage)
}

func (logger *Logger) Warn(messageFormat string, v ...interface{}) {
	formattedMessage := fmt.Sprintf(messageFormat, v...)
	_, file, line, _ := runtime.Caller(1)
	fmt.Fprintf(logger.warnWriter, "[%s][WARN][" + logger.yourType.Name() + "][%s:%d] : %s\n", time.Now().Format(time.RFC3339), file, line, formattedMessage)
}

func (logger *Logger) Error(messageFormat string, v ...interface{}) {
	formattedMessage := fmt.Sprintf(messageFormat, v...)
	_, file, line, _ := runtime.Caller(1)	
	fmt.Fprintf(logger.errorWriter, "[%s][ERROR][" + logger.yourType.Name() + "][%s:%d] : %s\n", time.Now().Format(time.RFC3339), file, line, formattedMessage)
}

func NewLogger(yourType reflect.Type, debugWriter io.Writer, infoWriter io.Writer, warnWriter io.Writer, errorWriter io.Writer) (*Logger) {
	var logger *Logger = new(Logger)
	logger.yourType = yourType
	logger.debugWriter = debugWriter
	logger.infoWriter = infoWriter
	logger.warnWriter = warnWriter
	logger.errorWriter = errorWriter
	return logger
}

func NewConsoleLogger(yourType reflect.Type) (*Logger) {
	if yourType == nil {
		panic("[NewConsoleLogger] Parameter `yourType` must not be nil")
	}

	return NewLogger(yourType, os.Stdout, os.Stdout, os.Stdout, os.Stderr)
}
