package main

import (
	"errors"

	"github.com/TheFredward/gator-cli-rss-feeder/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	command   string
	arguments []string
}

type commands struct {
	AllCmds map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Username is required to login")
	}
	userName := cmd.arguments[0]
	s.cfg.SetUser(userName)
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	value, ok := c.AllCmds[cmd.command]
	if !ok {
		return errors.New("Command not found")
	}
	err := value(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.AllCmds[name] = f
}
