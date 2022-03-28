package types

var GlobalConf GlobalConfMap

type GlobalConfMap struct {
	Global struct {
		Name       string `toml:"name"`
		Flag       bool   `toml:"flag"`
		TimeFormat string `toml:"time_format"`
		DateFormat string `toml:"date_format"`
	} `toml:"global"`
}
