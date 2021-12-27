package lext

import (
	"io"
	"log"
)

type InfoLogger interface {
	Info(f string, m ...interface{})
}

type WarnLogger interface {
	Warn(f string, m ...interface{})
}

type ErrorLogger interface {
	Error(f string, m ...interface{})
}

type ServerLogger interface {
	Server(f string, m ...interface{})
}

type KillLogger interface {
	Kill(f string, m ...interface{})
}

type Logger interface {
	InfoLogger
	WarnLogger
	ErrorLogger
	ServerLogger
	KillLogger
}

type PrettyLogger struct {
	info,
	warning,
	error,
	server *log.Logger
}

func New(out io.Writer) *PrettyLogger {
	return &PrettyLogger{
		info:    log.New(out, "\033[0;36minfo\033[0m => ", 0),
		warning: log.New(out, "\033[0;33mwarn\033[0m => ", 0),
		error:   log.New(out, "\033[0;31merror\033[0m => ", 0),
		server:  log.New(out, "\033[1;36mserver\033[0m => ", 0),
	}
}

func (l *PrettyLogger) Info(f string, m ...interface{}) {
	l.info.Printf(f, m...)
}

func (l *PrettyLogger) Warn(f string, m ...interface{}) {
	l.warning.Printf(f, m...)
}

func (l *PrettyLogger) Error(f string, m ...interface{}) {
	l.error.Printf(f, m...)
}

func (l *PrettyLogger) Server(f string, m ...interface{}) {
	l.server.Printf(f, m...)
}

func (l *PrettyLogger) Kill(f string, m ...interface{}) {
	l.error.Fatalf(f, m...)
}
