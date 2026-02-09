package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var partials []string = []string{
	"templates/base.layout.gohtml",
	"templates/head.partial.gohtml",
	"templates/footer.partial.gohtml",
}

//go:embed templates
var embedded embed.FS

const port = 80

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling request for:", r.URL.Path)
		render(w, "index.page.gohtml")
	})

	fmt.Println("Starting server on :", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func render(w http.ResponseWriter, specificTemplateName string) {
	var templateFileNames []string

	templateFileNames = append(templateFileNames, partials...)
	templateFileNames = append(templateFileNames, "templates/"+specificTemplateName)

	fullTemplate, err := template.ParseFS(embedded, templateFileNames...)
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
