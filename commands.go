package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/TheFredward/gator-cli-rss-feeder/internal/config"
	"github.com/TheFredward/gator-cli-rss-feeder/internal/database"
	"github.com/google/uuid"
)

type state struct {
	db  *database.Queries
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
	var checkUser database.User
	checkUser, err := s.db.GetUser(context.Background(), sql.NullString{String: userName, Valid: true})
	if err != nil {
		fmt.Println("User does not exist!")
		os.Exit(1)
	}
	s.cfg.SetUser(checkUser.Name.String)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("User name is required to register")
	}
	userName := cmd.arguments[0]
	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      sql.NullString{String: userName, Valid: true},
	}
	createdUser, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var foundUser database.User
	foundUser, err = s.db.GetUser(context.Background(), createdUser.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(foundUser)
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
