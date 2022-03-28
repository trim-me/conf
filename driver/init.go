package driver

import (
	"github.com/BurntSushi/toml"
	"log"
)

//init toml
func InitConfToml(file string, ConfMap interface{}) {
	if _, err := toml.DecodeFile(FileStat(file), ConfMap); err != nil {
		log.Fatal("init "+file+" failed:", err)
	}
}
