package fileFunctions

import (
	"fmt"
	"os"
)

type Config struct {
	Db_url   string `json:"db_url"`
	Curr_usr string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) {
	cfg.Curr_usr = userName
}

type state struct {
	configuration *Config
}

type command struct {
	command   string
	arguments []string
}

func handlerLogin(s *state, cmd command) {
	if len(cmd.arguments) == 0 {
		fmt.Print("Username is required to login")
		os.Exit(1)
	}
}
