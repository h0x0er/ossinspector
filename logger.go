package ossinspector

import "log"

type Logger struct {
	verbose bool
}

func (l *Logger) Printf(format string, v ...interface{}) {
	if l.verbose {
		log.Printf(format, v...)
	}
}
func (l *Logger) Println(v ...interface{}) {
	if l.verbose {
		log.Println(v...)
	}
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.verbose {
		log.Fatalf(format, v...)
	}
}

var logger *Logger

func NewLogger(verbose bool) *Logger {
	if logger == nil {
		logger = new(Logger)
		logger.verbose = verbose
	}
	return logger
}
