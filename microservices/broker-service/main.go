package main

import (
	"broker-service/internal/server"
	"fmt"
	"log"
	"net/http"
)

const port = 80

func main() {
	app := server.App{}

	log.Printf("Starting the Broker Service on port :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), app.Routes())
}
