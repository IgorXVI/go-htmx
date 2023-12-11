package main

import (
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Riddley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}

		tmpl.Execute(w, films)
	})

	http.HandleFunc("/add-film/", func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")

		director := r.PostFormValue("director")

		tmpl := template.Must(template.ParseFiles("index.html"))

		tmpl.ExecuteTemplate(w, "film-list-element", Film{
			Title:    title,
			Director: director,
		})
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
