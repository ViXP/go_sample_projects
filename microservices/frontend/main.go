package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var partials []string = []string{
	"./templates/base.layout.gohtml",
	"./templates/head.partial.gohtml",
	"./templates/footer.partial.gohtml",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling request for:", r.URL.Path)
		render(w, "index.page.gohtml")
	})

	fmt.Println("Starting server on :80")
	http.ListenAndServe(":80", nil)
}

func render(w http.ResponseWriter, specificTemplateName string) {
	var templateFileNames []string

	templateFileNames = append(templateFileNames, partials...)
	templateFileNames = append(templateFileNames, "./templates/"+specificTemplateName)

	fullTemplate, err := template.ParseFiles(templateFileNames...)
	if err != nil {
		log.Panic(err)
	}

	if err := fullTemplate.Execute(w, templateData()); err != nil {
		log.Panic(err)
	}
}

func templateData() *map[string]string {
	return &map[string]string{
		"BrokerUrl": os.Getenv("BROKER_URL"),
		"HandleUrl": os.Getenv("BROKER_URL") + "/handle",
	}
}
