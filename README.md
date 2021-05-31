# load toml conf to init Mysql , Redis ... 
 toml is multi array with load 
- mysql(use gorm tool) cluster  init ,
- redis (use redisgo tool) cluster init
- simple log data


### Toml conf 
- if you  use the driver , you mush mkdir the "conf/api/xx.toml"  
- in "examples/conf/" content is specific  conf file to learn

### use example for Mysql 
```golang 
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
```

### use example for Redis 
```golang
func main() {
	//set variable
	var MasterConf driver.RedisConnConf
	//get conf
	MasterConf.InitRedisToml("/conf/api/redis.toml")
	//get redis pool by name
	MasterDb := MasterConf.GetPool("master")
	//get redis pool by name
	UserDb := MasterConf.GetPool("test")
	fmt.Println(MasterDb)
	fmt.Println(UserDb)
}
```

### use example for api conf  
```golang

func main() {
	//global variable's
	driver.GlobalConf.InitGlobalToml("/conf/global.toml")
	fmt.Println(driver.GlobalConf.Global)
	//get api conf
	driver.ConfMap.InitConfToml("/conf/api/conf.toml")
	fmt.Println(driver.ConfMap.Api)
}


```

###  Detailed example can be found in examples folders
```sh
$ cd ./examples
$ go run ./app/redis.go
output:
&{0x12c6ea0 <nil> <nil> 0 0 4m0s false 0s {0 0} false 0 {0 {0 0}} <nil> {0 <nil> <nil>} 0 0}
&{0x12c6ea0 <nil> <nil> 0 0 4m0s false 0s {0 0} false 0 {0 {0 0}} <nil> {0 <nil> <nil>} 0 0}

```