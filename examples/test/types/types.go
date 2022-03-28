package types

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
