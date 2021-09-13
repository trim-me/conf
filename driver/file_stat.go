package driver

import (
	"log"
	"os"
	"path/filepath"
)

//verify file isExist
func FileStat(file string) string {
	var (
		str, _          = os.Getwd()
		defaultConfPath = str + file
		execDir         string
		err             error
	)
	_, err = os.Stat(defaultConfPath)
	if os.IsNotExist(err) {
		if execDir, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
			log.Fatal("get exec file path failed")
		}
		defaultConfPath = execDir + file
		_, err = os.Stat(defaultConfPath)
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	return defaultConfPath
}