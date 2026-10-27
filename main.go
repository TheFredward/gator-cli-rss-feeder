package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/TheFredward/gator-cli-rss-feeder/internal/config/fileFunctions"
)

func main() {
	cfg := &fileFunctions.Config{}
	err := fileFunctions.Read(cfg)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	cfgState := &state{cfg: cfg}
	cmds := commands{
		AllCmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	userArgs := os.Args
	if len(userArgs) < 2 {
		fmt.Printf("%v\n", errors.New("Incorrect number of arguments! require 2"))
		os.Exit(1)
	}
	userCmd := userArgs[1]
	userCmdCall := command{
		command:   userCmd,
		arguments: userArgs[2:],
	}
	err = cmds.run(cfgState, userCmdCall)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Login successful. %s logged in\n", userArgs[2:])
	os.Exit(0)

}
