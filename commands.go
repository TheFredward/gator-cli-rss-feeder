package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func getUser(s *state, name string) (database.User, error) {
	user, err := s.db.GetUser(context.Background(), sql.NullString{String: name, Valid: true})
	if err != nil {
		return database.User{}, errors.New("User not found!")
	}
	return user, nil
}
func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Username is required to login")
	}
	userName := cmd.arguments[0]
	checkUser, err := getUser(s, userName)
	if err != nil {
		return err
	}
	s.cfg.SetUser(checkUser.Name.String)
	fmt.Printf("Login successful. %s logged in\n", userName)
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
	_, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return err
	}
	s.cfg.SetUser(userName)
	return nil

}

func handlerGetUsers(s *state, cmd command) error {
	listUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	if len(listUsers) == 0 {
		return errors.New("No users have been added yet!")
	}
	for _, user := range listUsers {
		if s.cfg.Curr_usr == user.String {
			fmt.Printf("%s (current)\n", user.String)
		} else {
			fmt.Println(user.String)
		}
	}
	return nil
}

func handlerFetchFeed(s *state, cmd command) error {
	data, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v", data)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) <= 1 {
		return errors.New("Require Title and link")
	}
	foundUser, err := getUser(s, s.cfg.Curr_usr)
	if err != nil {
		return err
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      sql.NullString{String: name, Valid: true},
		Url:       sql.NullString{String: url, Valid: true},
		UserID:    uuid.NullUUID{UUID: foundUser.ID, Valid: true},
	}
	feed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", feed)
	return nil
}

func handlerGetFeed(s *state, cmd command) error {
	data, err := s.db.GetFeed(context.Background())
	if err != nil {
		return err
	}
	for _, feedVal := range data {
		user, err := s.db.GetUserUuid(context.Background(), feedVal.UserID.UUID)
		if err != nil {
			return err
		} else {
			fmt.Printf("Name: %s\nURL: %s\nUser: %s\n", feedVal.Name.String, feedVal.Url.String, user.Name.String)
		}

	}
	return nil
}

func reset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}
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
