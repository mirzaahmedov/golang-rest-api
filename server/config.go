package server

import (
	"api/store"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	PORT string
	Path string
	Store *store.Config
}

func NewConfig(path string) *Config {
	return &Config{
		Path: path,
		Store: store.NewConfig(),
	}
}

func (c *Config)Load() error {
	if _, err := toml.DecodeFile(c.Path, c); err != nil {
		log.Fatal(err)
	}
	return nil
}