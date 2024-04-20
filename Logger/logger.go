package Logger

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

var globalLogger *Logger = nil

func newLogger() *Logger {
	return &Logger{log.New(os.Stdout, "[revw-log] ", log.Ldate|log.Ltime)}
}

func GetLogger() *log.Logger {
	if globalLogger == nil {
		globalLogger = newLogger()
	}

	return globalLogger.logger
}
