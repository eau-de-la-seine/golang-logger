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
        var log *logger.Logger = logger.NewConsoleLogger(yourType, 0)
        log.Error("Good morning Agent %d", 47)
        log.Error("Good morning Agent %d", 48)
    }

    // Print result
    [2018-10-19T16:34:43+02:00][ERROR][YourType][/home/gekinci/projects/golang-logger/main.go:13] : Good morning Agent 47
    [2018-10-19T16:34:43+02:00][ERROR][YourType][/home/gekinci/projects/golang-logger/main.go:14] : Good morning Agent 48

