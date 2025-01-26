package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	mu       sync.Mutex
	level    string
	output   *os.File
	logLevel map[string]int
}

func New(level string) *Logger {
	logLevel := map[string]int{
		"DEBUG": 1,
		"INFO":  2,
		"WARN":  3,
		"ERROR": 4,
	}
	return &Logger{
		level:    level,
		output:   os.Stdout,
		logLevel: logLevel,
	}
}

func (l *Logger) log(level string, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.logLevel[level] < l.logLevel[l.level] {
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formatted := fmt.Sprintf("%s [%s] %s\n", timestamp, level, msg)
	_, _ = l.output.WriteString(formatted)
}

func (l *Logger) Info(msg string) {
	l.log("INFO", msg)
}

func (l *Logger) Error(msg string) {
	l.log("ERROR", msg)
}

func (l *Logger) Debug(msg string) {
	l.log("DEBUG", msg)
}

func (l *Logger) Warn(msg string) {
	l.log("WARN", msg)
}
