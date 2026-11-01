package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/TheFredward/gator-cli-rss-feeder/internal/config"
	"github.com/TheFredward/gator-cli-rss-feeder/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg := &config.Config{}
	err := config.Read(cfg)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
		return
	}
	var db *sql.DB
	db, err = sql.Open("postgres", cfg.Db_url)
	if err != nil {
		fmt.Print(err)
	}
	dbQueries := database.New(db)
	cfgState := &state{cfg: cfg, db: dbQueries}
	cmds := commands{
		AllCmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
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
