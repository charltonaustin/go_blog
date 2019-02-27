package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type blogPost struct {
	Single      string
	Name        string
	PublishDate string
	Content     template.HTML
}

type archiveLinks struct {
	LinkHref string
	LinkDate string

}
func main() {
	log.Printf("Listening on %v; ctrl + c to stop", ":9000")
	router := mux.NewRouter()
	md := []byte("## markdown document")
	output := markdown.ToHTML(md, nil, nil)

	router.Methods("GET").Path("/{year}/{month}/{day}/{name}").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		yearStr := vars["year"]
		monthStr := vars["month"]
		dayStr := vars["day"]
		name := vars["name"]
		content, err := ioutil.ReadFile(fmt.Sprintf("/Users/charltonaustin/dev/personal/blog-entries/published/%v/%v/%v/%v.md", yearStr, monthStr, dayStr, name))
		if err != nil {
			log.Printf("in error condition")
			NotFound(w, r)
			return
		}
		output := markdown.ToHTML(content, nil, nil)
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/home.html", "templates/blog.html", "templates/sidebar.html", "templates/pager.html"))
		tmpl.Execute(w, struct {
			DescriptionTag string
			Content        string
			TitleTag       string
			HomeActive     string
			AboutActive    string
			BlogPosts      []blogPost
			Previous       string
			Next           string
			ArchiveLinks   []archiveLinks
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
				Content:     template.HTML(output),
			}},
			Previous:     "/",
			Next:         "2018/02",
			ArchiveLinks: []archiveLinks{{LinkDate: "Feb 2018", LinkHref: "2018/02"}},
		})

	}))
	router.Methods("GET").Path("/{year}/{month}")


	router.Methods("GET").Path("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(
			"templates/base.html",
			"templates/home.html",
			"templates/blog.html",
			"templates/sidebar.html",
			"templates/pager.html",
		))
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
				Content:     template.HTML(output),
			}},
			Previous:     "/",
			Next:         "2018/02",
			ArchiveLinks: []archiveLinks{{LinkDate: "Feb 2018", LinkHref: "2018/02"}},
		})
	}))

	router.Methods("GET").Path("/about").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("/Users/charltonaustin/dev/personal/blog-entries/about.md")
		if err != nil {
			panic("unable to read about file")
		}
		output := markdown.ToHTML(content, nil, nil)
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/about.html"))
		tmpl.Execute(w, struct {
			DescriptionTag string
			TitleTag       string
			Content        template.HTML
			HomeActive     string
			AboutActive    string
		}{
			DescriptionTag: "description tag",
			Content:        template.HTML(output),
			TitleTag:       "This is a title",
			HomeActive:     "",
			AboutActive:    "active",

		})
	}))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/public"))))
	router.NotFoundHandler = http.HandlerFunc(NotFound)
	http.ListenAndServe(":9000", router)
}
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("in not found handler")
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	tmpl.Execute(w, struct {
		Title   string
		Message string
		Status  string
	}{
		Title:   "Page not found",
		Message: "We could not find the page you were looking for",
		Status:  "404",
	})
}