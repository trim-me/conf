package src

import (
	"log"
	"os"
)

func FileStat(file string) string {
	var (
		str, _          = os.Getwd()
		defaultConfPath = str + file
	)
	_, err := os.Stat(defaultConfPath)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}
	return defaultConfPath
}
