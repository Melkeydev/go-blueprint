package dbdriver

type SqliteTemplate struct{}

func (m SqliteTemplate) Service() []byte {
	return []byte(`package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Service interface {
	Health() map[string]string
}

type service struct {
	db *sql.DB
}

func New() *service {
	const file string = "example.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
`)
}
