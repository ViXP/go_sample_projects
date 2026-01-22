package main

import (
	"log"
	"logger-service/internal/data"
	"logger-service/internal/server"
	"os"
)

const (
	port    = 80
	rpcPort = 5001
)

func main() {
	client, err := server.InitializeDBConnection(os.Getenv("MONGODB_URL"))
	if err != nil {
		log.Panic(err)
	}

	store := data.NewStore(client)

	server.InitializeServer(port, store)
}
