package golog

import (
	"log"
	"os"
)

var (
	LogFile *os.File
	Logger  *log.Logger
)

func init() {
	InitLog()
}

func InitLog() {

	/* 创建日志文件 */
	LogFile, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err.Error())
	}
	Logger = log.New(LogFile, "GoLesson", log.Lshortfile|log.Ltime)

}
