package main

import (
	"fmt"

	"github.com/TheFredward/gator-cli-rss-feeder/internal/config/fileFunctions"
)

func main() {
	cfg := &fileFunctions.Config{}
	err := fileFunctions.Read(cfg)
	if err != nil {
		fmt.Println(err)
	} else {
		cfg.SetUser("freddie")
		fileFunctions.Write(cfg)
	}
	fmt.Println("Re-reading after updating user")
	fileFunctions.Read(cfg)
	if err != nil {
		fmt.Println(err)
	}
}
