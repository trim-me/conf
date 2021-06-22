package main

import (
	"fmt"
	"github.com/trim-me/conf/driver"
)

func main() {
	//set variable
	var MasterConf driver.RedisConnConf
	//get conf
	var path = "/conf/api/redis.toml"
	MasterConf.InitRedisToml(path)
	//get redis pool by name
	MasterDb := MasterConf.GetPool("master")
	//get redis pool by name
	UserDb := MasterConf.GetPool("test")
	fmt.Println(MasterDb)
	fmt.Println(UserDb)
}
