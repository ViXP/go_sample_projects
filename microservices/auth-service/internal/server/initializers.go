package server

import (
	"auth-service/internal/controllers"
	"auth-service/internal/data"
	"auth-service/internal/grpc_server"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

const (
	port     = 80
	grpcPort = 50051
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

func InitializeServer(store *data.Store) error {
	log.Println("Starting authentication REST server")

	return http.ListenAndServe(fmt.Sprintf(":%v", port), Routes(store))
}

func InitializeGRPCServer(store *data.Store) error {
	var protocols http.Protocols
	protocols.SetUnencryptedHTTP2(true)

	s := grpc.NewServer()

	grpc_server.RegisterAuthServiceServer(s, controllers.NewUsersGRPCController(store))

	http2Handler := h2c.NewHandler(s, &http2.Server{})

	server := http.Server{
		Addr:      fmt.Sprintf(":%v", grpcPort),
		Handler:   http2Handler,
		Protocols: &protocols,
	}

	defer server.Close()

	log.Println("Starting authentication gRPC server")
	return server.ListenAndServe()
}
