package config

type Config struct {
	Db_url   string `json:"db_url"`
	Curr_usr string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) {
	cfg.Curr_usr = userName
	Write(cfg)
}
