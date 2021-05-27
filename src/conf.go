package src

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

//init toml
func InitToml(file string)(conf interface{})  {
	var (
		str, _          = os.Getwd()
		defaultConfPath = str + file
	)
	_, err := os.Stat(defaultConfPath)
	if os.IsNotExist(err) {
		log.Fatal(err)
	}
	if _, err := toml.DecodeFile(defaultConfPath, &conf); err != nil {
		log.Fatal("init "+file+" failed:", err)
	}
	return
}
