package main

import (
	"fmt"
	"github.com/trim-me/conf/driver"
	"github.com/trim-me/conf/examples/test/types"
)

var (
	ConfMap    types.AppConfMap
	GlobalConf types.GlobalConfMap
)

func main() {
	//get api conf
	AppConfToml()

	//get global conf
	GlobalConfToml()

}
func AppConfToml() {
	AppConfPath := "/conf/app/conf.toml"
	//get app conf
	driver.InitConfToml(AppConfPath, &ConfMap)
	fmt.Println(ConfMap.App.HttpPort)
}

func GlobalConfToml() {
	GlobalConfPath := "/conf/global.toml"
	//get api conf
	driver.InitConfToml(GlobalConfPath, &GlobalConf)
	fmt.Println(GlobalConf)
}
