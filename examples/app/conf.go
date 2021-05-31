package main

import (
	"fmt"
	"github.com/trim-me/conf/driver"
)

func main() {
	//global variable's
	driver.GlobalConf.InitGlobalToml("/conf/global.toml")
	fmt.Println(driver.GlobalConf.Global)
	//get api conf
	driver.ConfMap.InitConfToml("/conf/api/conf.toml")
	fmt.Println(driver.ConfMap.Api)
}
