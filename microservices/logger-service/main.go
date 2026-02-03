package main

import (
	"log"
	"logger-service/internal/data"
	"logger-service/internal/server"
	"os"
)

const (
	port = 80
)

func main() {
	client, disconnect, err := server.InitializeDBConnection(os.Getenv("MONGODB_URL"))
	if err != nil {
		log.Panic(err)
	}

	defer disconnect()

	store := data.NewStore(client)

	server.InitializeServer(port, store)
}
