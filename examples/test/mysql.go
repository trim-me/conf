package main

import (
	"fmt"
	"github.com/trim-me/conf/driver"
)

func main() {
	//set variable
	var MysqlDriver driver.MysqlConnConf
	//get conf
	var path = "/conf/app/mysql.toml"
	MysqlDriver.InitMysqlToml(path)
	//get db pool by name
	MasterDb := MysqlDriver.GetPool("master")
	//get db pool by name
	UserDb := MysqlDriver.GetPool("user")
	fmt.Println(MasterDb)
	fmt.Println(UserDb)
}
