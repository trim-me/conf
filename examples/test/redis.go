package main

import (
	"fmt"
	"github.com/trim-me/conf/driver"
)

func main() {
	//set variable
	var RedisDriver driver.RedisConnConf
	//get conf
	var path = "/conf/app/redis.toml"
	RedisDriver.InitRedisToml(path)
	//get redis pool by name
	MasterDb := RedisDriver.GetPool("master")
	//get redis pool by name
	UserDb := RedisDriver.GetPool("test")
	fmt.Println(MasterDb)
	fmt.Println(UserDb)
}
