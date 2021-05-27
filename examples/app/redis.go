package main

import (
	"driver/src"
	"fmt"
)

func main() {
	//set variable
	var MasterConf src.RedisConnConf
	//get conf
	MasterConf.InitRedisToml("/conf/api/redis.toml")
	//get redis pool by name
	MasterDb := MasterConf.GetPool("master")
	//get redis pool by name
	UserDb := MasterConf.GetPool("test")
	fmt.Println(MasterDb)
	fmt.Println(UserDb)
}
