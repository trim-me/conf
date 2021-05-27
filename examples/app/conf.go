package main

import (
	"driver/src"
	"fmt"
)

func main() {
	src.GetGlobal()
	GetNameConf()
}

//get conf by name
func GetNameConf() {
	var conf src.ApiConf
	conf.InitConfToml("/conf/api/conf.toml")
	ApiConf := src.ConfMap.Conf["api"]
	fmt.Println(ApiConf)
}

