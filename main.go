package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/example.html"))
		tmpl.Execute(w, nil)
	})

	log.Print("Listeting port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
