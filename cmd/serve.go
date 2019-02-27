package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

type BlogPostPaths struct {
	year  string
	month string
	day   string
	name  string
}
func main() {
	router := mux.NewRouter()

	router.Methods("GET").Path("/{year}/{month}/{day}/{name}").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		yearStr := vars["year"]
		monthStr := vars["month"]
		dayStr := vars["day"]
		name := vars["name"]
		content, err := ioutil.ReadFile(fmt.Sprintf("/Users/charltonaustin/dev/personal/blog-entries/published/%v/%v/%v/%v.md", yearStr, monthStr, dayStr, name))
		if err != nil {
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
		var paths []BlogPostPaths
		filepath.Walk(
			"/Users/charltonaustin/dev/personal/blog-entries/published",
			filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() && strings.Contains(info.Name(), ".md") {
					path = strings.Replace(path, "/Users/charltonaustin/dev/personal/blog-entries/published/", "", -1)
					date := strings.Split(path, "/")
					paths = append(
						paths,
						BlogPostPaths{
							year:  date[0],
							month: date[1],
							day:   date[2],
							name:  info.Name(),
						},
					)
				}
				return nil
			}))

		var blogPosts []blogPost
		for i := len(paths) - 1; i > len(paths)-4; i-- {
			blogPostInfo := paths[i]
			content, err := ioutil.ReadFile(fmt.Sprintf("/Users/charltonaustin/dev/personal/blog-entries/published/%v/%v/%v/%v", blogPostInfo.year, blogPostInfo.month, blogPostInfo.day, blogPostInfo.name))
			if err != nil {
				panic(err)
			}
			blogPosts = append(blogPosts, blogPost{
				Single:      "",
				Name:        blogPostInfo.name,
				PublishDate: fmt.Sprintf("%v/%v/%v", blogPostInfo.year, blogPostInfo.month, blogPostInfo.day),
				Content:     template.HTML(content),
			})
		}

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
			DescriptionTag: "latest three blog posts",
			Content:        "This is the blog page",
			TitleTag:       "This is a title",
			HomeActive:     "active",
			AboutActive:    "",
			BlogPosts:      blogPosts,
			Previous:       "/",
			Next:           "2018/02",
			ArchiveLinks:   []archiveLinks{{LinkDate: "Feb 2018", LinkHref: "2018/02"}},
		})
	}))

	router.Methods("GET").Path("/about").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("/Users/charltonaustin/dev/personal/blog-entries/about.md")
		if err != nil {
			panic(err)
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
	log.Printf("Listening on %v; ctrl + c to stop", ":9000")
	http.ListenAndServe(":9000", router)
}
func NotFound(w http.ResponseWriter, r *http.Request) {
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