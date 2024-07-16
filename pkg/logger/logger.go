package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
}

type ColorLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func NewLogger() Logger {
	return &ColorLogger{
		infoLogger:  log.New(os.Stdout, "", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "", log.Ldate|log.Ltime),
	}
}

func formatMessage(msg string, keysAndValues ...interface{}) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%-20s", msg))

	for i := 0; i < len(keysAndValues); i += 2 {
		key := keysAndValues[i]
		value := keysAndValues[i+1]

		// Форматирование чисел с плавающей точкой
		if v, ok := value.(float64); ok {
			sb.WriteString(fmt.Sprintf("| %v: %.2f ", key, v))
		} else {
			sb.WriteString(fmt.Sprintf("| %v: %v ", key, value))
		}
	}

	return sb.String()
}

func (l *ColorLogger) Info(msg string, keysAndValues ...interface{}) {
	l.infoLogger.Printf("%sINFO%s  | %s%s", colorGreen, colorReset, formatMessage(msg, keysAndValues...), colorReset)
}

func (l *ColorLogger) Error(msg string, keysAndValues ...interface{}) {
	l.errorLogger.Printf("%sERROR%s | %s%s", colorRed, colorReset, formatMessage(msg, keysAndValues...), colorReset)
}

func (l *ColorLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.errorLogger.Fatalf("%sFATAL%s | %s%s", colorPurple, colorReset, formatMessage(msg, keysAndValues...), colorReset)
}
