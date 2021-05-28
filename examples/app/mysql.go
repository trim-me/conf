package main

import (
	"fmt"
	"github.com/trim-me/conf/driver"
)

func main() {
	//set variable
	var MasterConf driver.MysqlConnConf
	//get conf
	MasterConf.InitMysqlToml("/conf/api/mysql.toml")
	//get db pool by name
	MasterDb := MasterConf.GetPool("master")
	//get db pool by name
	UserDb := MasterConf.GetPool("user")
	fmt.Println(MasterDb)
	fmt.Println(UserDb)
}
