package src

import (
	"log"
	"os"
)

//初始化日志文件
func Loginit(name string) *log.Logger {
	file := "../logs/" + name + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal("初始化日志失败:", err)
	}
	return log.New(logFile, "["+name+"]", log.Ldate|log.Ltime|log.Lshortfile)
}
