package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main() {
	log.Printf("Listening on %v; ctrl + c to stop", ":9000")
	router := mux.NewRouter()
	router.Methods("GET").Path("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/base.html"))
		tmpl.Execute(w, struct{
			DescriptionTag string
			Content string
			TitleTag string
			HomeActive string
		}{
			DescriptionTag: "description tag",
			Content: "Hello, world!",
			TitleTag: "This is a title",
		})
	}))

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("resources/public")))

	http.ListenAndServe(":9000", router)
}
