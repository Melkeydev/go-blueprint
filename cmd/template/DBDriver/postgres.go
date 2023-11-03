package DBDriver

type PostgresTemplate struct{}

func (m PostgresTemplate) Service() []byte {
	return []byte(`package services
import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Service struct {
	db *sql.DB
}

func New() *Service {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
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
