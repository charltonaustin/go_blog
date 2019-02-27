package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

type blogPost struct {
	Single      string
	Name        string
	PublishDate string
	Content     string
}

type archiveLinks struct {

}
func main() {
	log.Printf("Listening on %v; ctrl + c to stop", ":9000")
	router := mux.NewRouter()
	router.Methods("GET").Path("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/home.html", "templates/blog.html",  "templates/sidebar.html", "templates/pager.html"))
		tmpl.Execute(w, struct {
			DescriptionTag string
			Content        string
			TitleTag       string
			HomeActive     string
			AboutActive    string
			BlogPosts      []blogPost
			Previous		string
			Next			string
			ArchiveLinks	[]archiveLinks
		}{
			DescriptionTag: "description tag",
			Content:        "This is the blog page",
			TitleTag:       "This is a title",
			HomeActive:     "active",
			AboutActive:    "",
			BlogPosts: []blogPost{{
				Single:      "single-",
				Name:        "post name",
				PublishDate: time.Now().Format("2006-01-02"),
				Content:     "post content",
			}},
			Previous: "previous",
			Next: "next",
		})
	}))

	router.Methods("GET").Path("/about").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/about.html"))
		tmpl.Execute(w, struct {
			DescriptionTag string
			Content        string
			TitleTag       string
			HomeActive     string
			AboutActive    string
		}{
			DescriptionTag: "description tag",
			Content:        "This is the about page",
			TitleTag:       "This is a title",
			HomeActive:     "active",
			AboutActive:    "",

		})
	}))

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("resources/public")))

	http.ListenAndServe(":9000", router)
}
