# A lightweight logger for Golang

## Specificities

This lightweight logger comes with some specificities you need to know :

* Each logger instance is unique, which means each call to `NewLogger` will give you a new object. Your logger instance **is threadsafe but** you have to be careful in a multithreaded environment if you have multiple logger instances trying to write into the same resource (the same file for example), because the logger instances do not share the same lock. A solution is to inject your logger instance.
* The message format is : `[<RFC3339-date>][<logger-level>][<file-name>.<file-line>] : <your-message>`
* Experimental `LoggerFactory` added for getting the same logger instance.


## Constructors

    NewLogger(level int, writer io.Writer) *Logger
    NewConsoleLogger(level int) *Logger



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
