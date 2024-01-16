package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/leoff00/go-marvel-api/api"
)

type ResponseInfo struct {
	Desc, Url, Alternative string
}

func handler1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html", "./image.html"))
	if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Println("Error during parse", err)
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received from HTMX")
	var responseInfo ResponseInfo

	formValue := r.PostFormValue("superhero")
	marvel := api.DoRequest(formValue)

	for _, v := range marvel.Data.Results {
		responseInfo.Desc = v.Description
		responseInfo.Url = fmt.Sprintf("%s.%s", v.Thumbnail.Path, v.Thumbnail.Extension)
		responseInfo.Alternative = "A Hero"
	}

	tmpl := template.Must(template.ParseFiles("index.html", "image.html"))
	if err := tmpl.ExecuteTemplate(w, "response_data", responseInfo); err != nil {
		log.Println("Cannot load the response info data", err)
	}
}

func main() {
	log.Default().Println("Listening on PORT http://localhost:3000")
	http.HandleFunc("/", handler1)
	http.HandleFunc("/search", handler2)
	http.ListenAndServe(":3000", nil)
}
