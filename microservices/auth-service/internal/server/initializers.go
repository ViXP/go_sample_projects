package server

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func InitializeDBConnection() (*sql.DB, error) {
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

func InitializeServer(app *App, port int) error {
	log.Println("Starting authentication service")

	return http.ListenAndServe(fmt.Sprintf(":%v", port), app.Routes())
}
