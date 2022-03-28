package main

import (
	"fmt"
	"github.com/trim-me/conf/driver"
	"github.com/trim-me/conf/examples/test/types"
)

var ConfMap AppConfMap

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

func main() {
	var AppConfPath = "/conf/app/conf.toml"
	//get api conf
	driver.InitConfToml(AppConfPath, &ConfMap)
	fmt.Println(ConfMap)

	//get global conf
	GlobalConfToml()

}

func GlobalConfToml() {
	var GlobalConfPath = "/conf/global.toml"
	//get api conf
	driver.InitConfToml(GlobalConfPath, &types.GlobalConf)
	fmt.Println(types.GlobalConf)
}
