package src

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

var ConfMap ApiConfMap
var GlobalConf GlobalConfMap

//init toml
func (conf *ApiConf) InitConfToml(file string) {
	if _, err := toml.DecodeFile(FileStat(file), &ConfMap); err != nil {
		log.Fatal("init "+file+" failed:", err)
	}
}

func GetApi() (Api *ApiConf) {
	var conf ApiConf
	conf.InitConfToml("/conf/api/conf.toml")
	Api = ConfMap.Conf["api"]
	return
}

//init api toml
func (conf *Global) InitGlobalToml(file string) {
	if _, err := toml.DecodeFile(FileStat(file), &GlobalConf); err != nil {
		log.Fatal("init "+file+" failed:", err)
	}
	fmt.Println(GlobalConf.Global)
}

//init global toml
func GetGlobal()  {
	var conf Global
	conf.InitGlobalToml("/conf/global.toml")
}
