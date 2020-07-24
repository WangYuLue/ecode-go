package log

import (
	"fmt"
	"log"
	"os"
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
	setPrefix(Level.Debug, 0)
	logger.Println(v...)
}

// Info -
func Info(v ...interface{}) {
	setPrefix(Level.Info, 0)
	logger.Println(v...)
}

// Warn -
func Warn(v ...interface{}) {
	setPrefix(Level.Warn, 0)
	logger.Println(v...)
}

// Error -
func Error(v ...interface{}) {
	setPrefix(Level.Error, 0)
	logger.Println(v...)
}

// Fatal -
func Fatal(v ...interface{}) {
	setPrefix(Level.Fatal, 0)
	logger.Fatalln(v...)
}

// ErrorWithDeep -
func ErrorWithDeep(deep int, v ...interface{}) {
	setPrefix(Level.Error, deep)
	logger.Println(v...)
}

func setPrefix(level string, deep int) {
	pc, file, line, ok := runtime.Caller(defaultCallerDepth + deep)
	name := runtime.FuncForPC(pc).Name()
	// fileList := strings.Split(file, "ecode-go")
	// if len(fileList) > 1 {
	// 	file = fileList[1]
	// }
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d][%s]\r", level, file, line, name)
	} else {
		logPrefix = fmt.Sprintf("[%s]\r", level)
	}
	logger.SetPrefix(logPrefix)
}
