package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.config.DatabaseHost,
		s.config.DatabasePort,
		s.config.DatabaseUser,
		s.config.DatabasePassword,
		s.config.DatabaseName,
	))

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}
