package logger

import (
	"fmt"
	"log"
)

type Logger struct {
}

func (l *Logger) Info(msg string, args ...interface{}) {
	log.Println(fmt.Sprintf(" [INFO] %s", fmt.Sprintf(msg, args...)))
}

func (l *Logger) Panic(err error, msg string, args ...interface{}) {
	log.Println(fmt.Sprintf(" [ERROR] %s - ERROR: %s", fmt.Sprintf(msg, args...), err))

	panic(err)
}

// Static functions

func NewLogger() *Logger {
	return &Logger{}
}
