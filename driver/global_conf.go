package driver

import (
	"github.com/BurntSushi/toml"
	"log"
)

var GlobalConf GlobalConfMap

type GlobalConfMap struct {
	Global *Global `toml:"global"`
}

//init api toml
func (conf *GlobalConfMap) InitGlobalToml(file string) {
	if _, err := toml.DecodeFile(FileStat(file), &GlobalConf); err != nil {
		log.Fatal("init "+file+" failed:", err)
	}
}
