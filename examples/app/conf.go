package main

import (
	"fmt"
	"github.com/trim-me/conf/driver"
)

func main() {
	//global variable's
	driver.GetGlobal()
	//get conf with specific filename
	GetNameConf()
}

//get conf by name
func GetNameConf() {
	var conf driver.ApiConf
	conf.InitConfToml("/conf/api/conf.toml")
	ApiConf := driver.ConfMap.Conf["api"]
	fmt.Println(ApiConf)
}

