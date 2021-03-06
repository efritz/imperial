package base

import "log"

type (
	// Logger is an interface to the logger the reporter writes to.
	Logger interface {
		// Prints logs a message. Arguments should be handled in the manner of fmt.Printf.
		Printf(format string, args ...interface{})
	}

	printLogger struct{}
	nilLogger   struct{}
)

// NilLogger is a singleton silent logger.
var NilLogger = NewNilLogger()

// NewPrintLogger creates a logger that prints to stdout.
func NewPrintLogger() Logger {
	return &printLogger{}
}

// NewNilLogger creates a silent logger.
func NewNilLogger() Logger {
	return &nilLogger{}
}

func (l *printLogger) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *nilLogger) Printf(format string, args ...interface{}) {
}
