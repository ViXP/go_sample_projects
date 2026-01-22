package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (le *LogEntry) collection() *mongo.Collection {
	return dbPool.Database("logs").Collection("logs")
}

func (le *LogEntry) Create(entry *LogEntry) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	timeStamp := time.Now()

	_, err := le.collection().InsertOne(ctx, LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: timeStamp,
		UpdatedAt: timeStamp,
	})
	return err
}

func (le *LogEntry) Find(id string) (*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	result := le.collection().FindOne(ctx, bson.M{"_id": id})

	var entry LogEntry

	err := result.Decode(&entry)

	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (le *LogEntry) Save() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	return le.collection().UpdateByID(ctx, le.ID, bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: le.Name},
		{Key: "data", Value: le.Data},
		{Key: "updated_at", Value: time.Now()},
	}}})
}

func (le *LogEntry) Destroy() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	result, err := le.collection().DeleteOne(ctx, bson.M{"_id": le.ID})

	if err != nil {
		return false, err
	}

	if result.DeletedCount > 0 {
		return true, nil
	}

	return false, nil
}

func (le *LogEntry) All() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	queryOptions := options.Find().SetSort(bson.D{{Key: "_id", Value: "1"}})

	cursor, err := le.collection().Find(ctx, bson.D{}, queryOptions)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var entries []*LogEntry

	if err = cursor.All(ctx, &entries); err != nil {
		return nil, err
	}

	return entries, nil
}
