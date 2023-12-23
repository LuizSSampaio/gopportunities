package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	fatal   *log.Logger
	writer  io.Writer
}

func NewLogger() *Logger {
	writer := io.Writer(os.Stdout)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", log.Ldate|log.Ltime),
		info:    log.New(writer, "INFO: ", log.Ldate|log.Ltime),
		warning: log.New(writer, "WARNING: ", log.Ldate|log.Ltime),
		err:     log.New(writer, "ERROR: ", log.Ldate|log.Ltime),
		fatal:   log.New(writer, "FATAL: ", log.Ldate|log.Ltime),
		writer:  writer,
	}
}

// Debug Non formatted log
func (logger *Logger) Debug(v ...interface{}) {
	logger.debug.Println(v...)
}

// Debugf Formatted log
func (logger *Logger) Debugf(format string, v ...interface{}) {
	logger.debug.Printf(format, v...)
}

// Info Non formatted log
func (logger *Logger) Info(v ...interface{}) {
	logger.info.Println(v...)
}

// Infof Formatted log
func (logger *Logger) Infof(format string, v ...interface{}) {
	logger.info.Printf(format, v...)
}

// Warning Non formatted log
func (logger *Logger) Warning(v ...interface{}) {
	logger.warning.Println(v...)
}

// Warningf Formatted log
func (logger *Logger) Warningf(format string, v ...interface{}) {
	logger.warning.Printf(format, v...)
}

// Error Non formatted log
func (logger *Logger) Error(v ...interface{}) {
	logger.err.Println(v...)
}

// Errorf Formatted log
func (logger *Logger) Errorf(format string, v ...interface{}) {
	logger.err.Printf(format, v...)
}

// Fatal Non formatted log
func (logger *Logger) Fatal(v ...interface{}) {
	logger.fatal.Println(v...)
}

// Fatalf Formatted log
func (logger *Logger) Fatalf(format string, v ...interface{}) {
	logger.fatal.Printf(format, v...)
}
