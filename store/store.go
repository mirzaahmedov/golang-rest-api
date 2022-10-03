package store

import (
	"api/models"
	"database/sql"
	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db *sql.DB
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
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store)Close() {
	s.db.Close()
}

func (s *Store)RegisterUser(u *models.User) (*models.User, error) {
	if err := s.db.QueryRow(
		"INSERT INTO people (id, name) VALUES ($1,$2) RETURNING *", 
		u.Id, 
		u.Name,
	).Scan(&u.Id, &u.Name); err != nil {
		return nil, err
	}

	return u, nil
}