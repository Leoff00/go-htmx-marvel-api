package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/leoff00/go-marvel-api/api"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// query := r.URL.Query()
	// fmt.Println(query)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Request received from HTMX")
	formValue := r.PostFormValue("superhero")
	resp := api.DoRequest(formValue)

	w.Write(resp)
}

func main() {
	log.Default().Println("Listening on PORT http://localhost:3000")
	http.HandleFunc("/", handler1)
	http.HandleFunc("/search", handler2)
	http.ListenAndServe(":3000", nil)
}
