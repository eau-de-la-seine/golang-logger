# A Golang logger

## Constructors

    NewLogger(reflect.Type, writer io.Writer, level int) *Logger
    NewConsoleLogger(reflect.Type, level int) *Logger


## Method signatures

    Debug(messageFormat string, v ...interface{})
    Info(messageFormat string, v ...interface{})
    Warn(messageFormat string, v ...interface{})
    Error(messageFormat string, v ...interface{})


## NewConsoleLogger example

    package main

    import (
        "github.com/eau-de-la-seine/golang-logger"
        "reflect"
    )

    type YourType int

    func main() {
        var yourType reflect.Type = reflect.TypeOf((*YourType)(nil)).Elem()
        var log *logger.Logger = logger.NewConsoleLogger(yourType, logger.LEVEL_ERROR)
        log.Error("Good morning Agent %d", 47)
        log.Debug("Good morning Agent %d", 48)
    }

    // Print result
    [2018-10-19T16:34:43+02:00][ERROR][YourType][demo.go:13] : Good morning Agent 47
