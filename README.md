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
    var path = "/conf/app/mysql.toml"
    MasterConf.InitMysqlToml(path)
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
    var path = "/conf/app/redis.toml"
    MasterConf.InitRedisToml(path)
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

var (
	ConfMap    AppConfMap
	GlobalConf GlobalConfMap
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


type GlobalConfMap struct {
	Global struct {
		Name       string `toml:"name"`
		Flag       bool   `toml:"flag"`
		TimeFormat string `toml:"time_format"`
		DateFormat string `toml:"date_format"`
	} `toml:"global"`
}

type AppConfMap struct {
	App struct {
		HttpPort int    `toml:"http_port"`
		Mode     string `toml:"mode"`
		Limit    int    `toml:"limit"`
	} `toml:"app"`
	Domain struct {
		AdminDomain string   `toml:"admin_domain"`
		WebDomain   []string `toml:"web_domain"`
		ImageUrl    string   `toml:"image_url"`
	} `toml:"domain"`
	Token struct {
		TokenName   string `toml:"token_name"`
		TokenExpire int    `toml:"token_expire"`
	} `toml:"Token"`
	Path struct {
		LogPathRoot string `toml:"log_path_root"`
		PathRoot    string `toml:"path_root"`
	} `toml:"Path"`
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