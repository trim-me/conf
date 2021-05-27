package src

import "log"

type Config struct {
	Global *Global                  `json:"global"`
	Conf   map[string]ApiConf       `json:"time_format"`
	Mysql  map[string]MysqlConnConf `json:"mysql"`
	Redis  map[string]RedisConnConf `json:"redis"`
	Log    *Log                     `json:"log"`
}

type Log struct {
	Log *log.Logger
}

type Global struct {
	Name       string `json:"name"`
	Flag       string `json:"flag"`
	TimeFormat string `json:"time_format"`
}

type ApiConf struct {
	HttpPort    int
	Mode        string
	TokenName   string
	TokenExpire int
	WebDomain   []string
	AdminDomain string
	ImageUrl    string
	Limit       int
	LogPathRoot string
	PathRoot    string
}
