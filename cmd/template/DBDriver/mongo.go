package DBDriver

type MongoTemplate struct{}

func (m MongoTemplate) Service() []byte {
	return []byte(`package service

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	db *mongo.Database
}

func New(ctx context.Context) (*Service, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://yourIpAddress:yourPort"))
	if err != nil {
		return nil, err
	}

	return &Service{
		db: client.Database("example"),
	}, nil
}
`)
}
