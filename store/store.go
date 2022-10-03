package store

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Store struct {
	config *Config
}

func New() *Store {
	config := NewConfig()

	if _, err := toml.DecodeFile("configs/store.toml", config); err != nil {
		log.Fatal(err)
	}
	
	return &Store{
		config: config,
	}
}

func (s *Store)Open() error {
	return nil
}

func (s *Store)Close() {

}