package main

import (
	"auth-service/data"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const port = 80

type App struct {
	DB     *sql.DB
	Models *data.Models
}

func main() {
	pgConn, err := initializeDBConnection()

	if err != nil {
		log.Panic(err)
	}

	defer pgConn.Close()

	app := App{
		DB:     pgConn,
		Models: data.NewModels(pgConn),
	}

	initializeServer(&app)
}

func initializeDBConnection() (*sql.DB, error) {
	log.Println("Connecting to database")

	var pgConn *sql.DB
	var err error

	dbUrl := os.Getenv("POSTGRES_URL")

	if len(dbUrl) == 0 {
		return nil, errors.New("POSTGRES_URL environment variable is not set")
	}

	for {
		pgConn, err = sql.Open("pgx", dbUrl)

		if err == nil {
			break
		} else {
			log.Println("PostgreSQL is not ready")
			time.Sleep(1 * time.Second)
		}
	}

	return pgConn, nil
}

func initializeServer(app *App) {
	log.Println("Starting authentication service")

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), app.routes())

	if err != nil {
		log.Panic(err)
	}
}
