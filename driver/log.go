package driver

import (
	"log"
	"os"
)

type Err struct {
	err  error
	code int
	msg  string
}

//init run log of service
func Loginit(name string) *log.Logger {
	_ = os.MkdirAll("./log", os.ModePerm)
	file := "./log/" + name + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal("初始化日志失败:", err)
	}
	return log.New(logFile, "["+name+"]", log.Ldate|log.Ltime|log.Lshortfile)
}
