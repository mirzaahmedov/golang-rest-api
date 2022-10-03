package main

import (
	"api/server"
	"flag"
	"log"
)

var (
	configPath string
)

func init() {
	configPath = *flag.String("config", "configs/server.toml", "use custom config file")
}

func main() {
	config := server.NewConfig(configPath)

	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	api := server.New(config)

	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}