# A Golang logger

## Constructors

    NewLogger(reflect.Type, writer io.Writer, level int) *Logger
    NewConsoleLogger(reflect.Type, level int) *Logger


## Logger Levels

	LEVEL_ERROR
	LEVEL_WARN
	LEVEL_INFO
	LEVEL_DEBUG
	LEVEL_OFF


## Logger's method signatures

    Debug(messageFormat string, v ...interface{})
    Info(messageFormat string, v ...interface{})
    Warn(messageFormat string, v ...interface{})
    Error(messageFormat string, v ...interface{})


## NewConsoleLogger example

    package main

    import (
        "github.com/eau-de-la-seine/golang-logger"
    )

    func main() {
        var log *logger.Logger = logger.NewConsoleLogger(logger.LEVEL_ERROR)
        log.Error("Good morning Agent %d", 47)
        log.Debug("Good morning Agent %d", 48)
    }

    // Print result
    [2018-10-31T09:19:53+01:00][ERROR][demo.go:9] : Good morning Agent 47
