package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const fileName = "/.gatorconfig.json"

func Read(cfg *Config) error {
	var dataBytes []byte
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	fullPath := homeDir + fileName
	dataBytes, err = os.ReadFile(fullPath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(dataBytes, cfg); err != nil {
		return err
	}
	return nil
}
func Write(cfg *Config) error {
	var byteVal []byte
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	fullPath := homeDir + fileName
	byteVal, err = json.Marshal(cfg)
	if err != nil {
		fmt.Print("Error: ", err)
	}
	err = os.WriteFile(fullPath, byteVal, 0644)
	if err != nil {
		return errors.New("Unable to update username")
	}
	return nil
}
