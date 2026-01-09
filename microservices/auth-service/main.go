package main

import (
	"auth-service/internal/data"
	"auth-service/internal/server"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const port = 80

func main() {
	pgConn, err := server.InitializeDBConnection()

	if err != nil {
		log.Panic(err)
	}

	defer pgConn.Close()

	app := data.App{
		DB:     pgConn,
		Models: data.NewModels(pgConn),
	}

	err = server.InitializeServer(&app, port)

	if err != nil {
		log.Panic(err)
	}
}
