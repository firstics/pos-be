package driver

import (
	"log"
	"os"
	"time"

	"github.com/pos-be/pkg/config"
)

type Logger interface {
	Debug(msg string)
	Error(msg string)
	Info(msg string)
	Warn(msg string)
}

type logger struct {
	log *log.Logger
}

func NewLogger(cfg config.Config) Logger {
	log := log.New(os.Stdout, "", 0)
	return &logger{log}
}

func (l *logger) Debug(msg string) {
	l.log.Println(
		"DEBUG: "+msg,
		time.Now().Format(time.RFC3339),
	)
}

func (l *logger) Error(msg string) {
	l.log.Println(
		"ERROR: "+msg,
		time.Now().Format(time.RFC3339),
	)
}

func (l *logger) Info(msg string) {
	l.log.Println(
		"INFO: "+msg,
		time.Now().Format(time.RFC3339),
	)
}

func (l *logger) Warn(msg string) {
	l.log.Println(
		"WARN: "+msg,
		time.Now().Format(time.RFC3339),
	)
}

type NoopLogger struct{}

func (NoopLogger) Debug(msg string) {}
func (NoopLogger) Error(msg string) {}
func (NoopLogger) Info(msg string)  {}
func (NoopLogger) Warn(msg string)  {}

func NewNoopLogger() Logger {
	return &NoopLogger{}
}
