package driver

import (
	"github.com/BurntSushi/toml"
	"log"
)

var ConfMap ApiConfMap

type ApiConfMap struct {
	Api *ApiConf `toml:"api"`
}

//init toml
func (conf *ApiConfMap) InitConfToml(file string) {
	if _, err := toml.DecodeFile(FileStat(file), &ConfMap); err != nil {
		log.Fatal("init "+file+" failed:", err)
	}
}
