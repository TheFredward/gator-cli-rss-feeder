package fileFunctions

import (
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "/gatorconfig.json"

func Read(cfg *Config) error {
	var dataBytes []byte
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: ", err)
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
	fmt.Printf("%+v\n", cfg) // prints: &{X:1 Y:2}
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
	return nil
}
