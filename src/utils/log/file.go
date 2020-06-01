package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logSavePath = "runtime/logs/"
	logSaveName = "log"
	logFileExt  = "log"
	timeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", logSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s-%s.%s", logSaveName, time.Now().Format(timeFormat), logFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir(getLogFilePath())
	case os.IsPermission(err):
		log.Fatalf("无权限 :%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开文件失败 :%v", err)
	}
	return handle
}

func mkDir(path string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
