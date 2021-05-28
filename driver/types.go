package driver

import "log"

type ApiConfMap struct {
	Conf map[string]*ApiConf `toml:"conf"`
}

type ApiConf struct {
	HttpPort    int      `toml:"http_port"`
	Mode        string   `toml:"mode"`
	TokenName   string   `toml:"token_name"`
	TokenExpire int      `toml:"token_expire"`
	WebDomain   []string `toml:"web_domain"`
	AdminDomain string   `toml:"admin_domain"`
	ImageUrl    string   `toml:"image_url"`
	Limit       int      `toml:"limit"`
	LogPathRoot string   `toml:"log_path_root"`
	PathRoot    string   `toml:"path_root"`
}
type Log struct {
	Log *log.Logger
}

type GlobalConfMap struct {
	Global Global `toml:"global"`
}

type Global struct {
	Name       string `toml:"name"`
	Flag       bool   `toml:"flag"`
	TimeFormat string `toml:"time_format"`
}