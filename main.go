package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/leoff00/go-marvel-api/api"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	log.Println("Request received from HTMX")

	formValue := r.PostFormValue("superhero")
	marvel := api.DoRequest(formValue)
	var img string
	var ext string
	for _, v := range marvel.Data.Results {
		img, ext = v.Thumbail.Path, v.Thumbail.Extension
	}

	htmlstr := fmt.Sprintf("<img src='%s.%s' alt='Heroes' ", img, ext)
	tmpl, err := template.New("t").Parse(htmlstr)

	if err != nil {
		log.Println("Error during parse template image")
	}

	tmpl.Execute(w, nil)
}

func main() {
	log.Default().Println("Listening on PORT http://localhost:3000")
	http.HandleFunc("/", handler1)
	http.HandleFunc("/search", handler2)
	http.ListenAndServe(":3000", nil)
}
