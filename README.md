# A Golang logger

## Constructors

    NewLogger(reflect.Type, debug io.Writer, info io.Writer, warn io.Writer, err io.Writer) *Logger
    NewConsoleLogger(reflect.Type) *Logger


## Method signatures

    Debug(messageFormat string, v ...interface{})
    Info(messageFormat string, v ...interface{})
    Warn(messageFormat string, v ...interface{})
    Error(messageFormat string, v ...interface{})


## NewConsoleLogger example

    import (
        "github.com/eau-de-la-seine/golang-logger"
        "reflect"
    )

    func main() {
		var yourType reflect.Type = reflect.TypeOf((*YourType)(nil)).Elem()
		var log *Logger = NewConsoleLogger(yourType)
		log.Error("Good morning Agent %d", 47)
	}

    // Print result
    [2018-10-15T18:46:11+02:00][ERROR][YourType][/your/path/file.go:9] : Good morning Agent 47
