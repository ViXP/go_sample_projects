package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 80

type App struct{}

func main() {
	app := App{}

	log.Printf("Starting the Broker Service on port :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), app.Routes())
}
