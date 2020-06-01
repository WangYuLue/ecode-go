package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	f                  *os.File
	defaultPrefix      = ""
	defaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
)

// SLevel -
type SLevel struct {
	Debug string
	Info  string
	Warn  string
	Error string
	Fatal string
}

// Level -
var Level = SLevel{
	Debug: "DEBUG",
	Info:  "INFO",
	Warn:  "WARN",
	Error: "ERROR",
	Fatal: "FATAL",
}

func init() {
	filePath := getLogFileFullPath()
	f = openLogFile(filePath)
	logger = log.New(f, defaultPrefix, log.LstdFlags)
}

// Debug -
func Debug(v ...interface{}) {
	setPrefix(Level.Debug)
	logger.Println(v...)
}

// Info -
func Info(v ...interface{}) {
	setPrefix(Level.Info)
	logger.Println(v...)
}

// Warn -
func Warn(v ...interface{}) {
	setPrefix(Level.Warn)
	logger.Println(v...)
}

// Error -
func Error(v ...interface{}) {
	setPrefix(Level.Error)
	logger.Println(v...)
}

// Fatal -
func Fatal(v ...interface{}) {
	setPrefix(Level.Fatal)
	logger.Fatalln(v...)
}

func setPrefix(level string) {
	_, file, line, ok := runtime.Caller(defaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", level, filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", level)
	}
	logger.SetPrefix(logPrefix)
}
