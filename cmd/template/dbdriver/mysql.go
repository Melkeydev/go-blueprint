package dbdriver

type MysqlTemplate struct{}

func (m MysqlTemplate) Service() []byte {
	return []byte(`package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Service struct {
	db *sql.DB
}

func New() *Service {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

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
