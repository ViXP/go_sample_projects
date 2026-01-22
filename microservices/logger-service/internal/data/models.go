package data

import "go.mongodb.org/mongo-driver/v2/mongo"

var dbPool *mongo.Client

type Models struct {
	LogEntry *LogEntry
}

func NewModels(client *mongo.Client) *Models {
	dbPool = client

	return &Models{
		LogEntry: &LogEntry{},
	}
}
