package data

import "go.mongodb.org/mongo-driver/v2/mongo"

type Store struct {
	DB     *mongo.Client
	Models *Models
}

func NewStore(client *mongo.Client) *Store {
	return &Store{
		DB:     client,
		Models: NewModels(client),
	}
}
