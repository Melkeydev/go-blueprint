package DBDriver

type SqliteTemplate struct{}

func (m SqliteTemplate) Service() []byte {
	return []byte(`package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	db *sql.DB
}

func New() *Service {
	const file string = "example.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	s := &Service{db: db}
	return s
}

func (s *Service) Health() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("db down: %v", err))
		return
	}
}
`)
}
