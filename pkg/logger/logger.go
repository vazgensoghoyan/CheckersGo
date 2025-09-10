package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	info  *log.Logger
	error *log.Logger
	debug *log.Logger
}

func New() *Logger {
	return &Logger{
		info:  log.New(os.Stdout, "INFO:  ", log.LstdFlags),
		error: log.New(os.Stderr, "ERROR:  ", log.LstdFlags),
		debug: log.New(os.Stdout, "DEBUG:  ", log.LstdFlags),
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.info.Println(formatMessage(msg, args...))
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.error.Println(formatMessage(msg, args...))
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.debug.Println(formatMessage(msg, args...))
}

func formatMessage(msg string, args ...interface{}) string {
	if len(args) > 0 {
		return fmt.Sprintf(msg, args...)
	}
	return msg
}
