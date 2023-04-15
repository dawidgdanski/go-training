package logging

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

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
