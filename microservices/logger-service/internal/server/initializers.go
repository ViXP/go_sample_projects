package server

import (
	"context"
	"fmt"
	"log"
	"logger-service/internal/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitializeServer(port int, store *data.Store) error {
	log.Println("Starting Logger server...")

	return http.ListenAndServe(fmt.Sprintf(":%v", port), Routes(store))
}

func InitializeDBConnection(dbUrl string) (*mongo.Client, func(), error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(dbUrl).SetServerAPIOptions(serverAPI)

	log.Println("Connecting to Mongo database...")
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, nil, err
	}

	disconnect := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := client.Disconnect(ctx); err != nil {
			log.Println("Mongo disconnect error:", err)
		}
	}

	log.Println("Database is connected.")
	return client, disconnect, nil
}
