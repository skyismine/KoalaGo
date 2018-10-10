package logging

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Info(v ...interface{})
	Warnning(v ...interface{})
	Debug(v ...interface{})
	Error(v ...interface{})
}

type _Logger struct {
	syslogger *log.Logger
}

func (logger *_Logger) trace(level string, v ...interface{}) {
	logger.syslogger.Println(level, fmt.Sprint(v...))
}

func New(path string) Logger {
	file, err := os.Create(path)
	if err != nil {
		return nil
	}
	syslogger := log.New(file, "", log.LstdFlags|log.Lshortfile)
	return &_Logger{
		syslogger:syslogger,
	}
}

func (logger *_Logger) Info(v ...interface{}) {
	logger.trace("Info", v...)
}

func (logger *_Logger) Warnning(v ...interface{}) {
	logger.trace("Warnning", v...)
}

func (logger *_Logger) Debug(v ...interface{}) {
	logger.trace("Debug", v...)
}

func (logger *_Logger) Error(v ...interface{}) {
	logger.trace("Error", v...)
}