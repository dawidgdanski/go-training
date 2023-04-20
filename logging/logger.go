package logging

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

/*
The ConsoleLogger implements the Logger interface because of the method of
identical signature to the method in the interface definition.
*/
func (log ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

func Process(logger Logger) {
	logger.Log("Processed")
}

func ProcessViaConsole() {
	logger := ConsoleLogger{}
	Process(logger)
}
