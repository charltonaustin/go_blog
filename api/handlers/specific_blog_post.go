package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-blog/api"
	"go-blog/blog"
	"go-blog/interfaces"
	"log"
	"net/http"
)

func CreateSpecificBlogPostHandler(templateGetter api.TemplateGetter, pathGetter interfaces.PathGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SpecificBlogPost(w, r, templateGetter, pathGetter)
	})
}
func SpecificBlogPost(w http.ResponseWriter, r *http.Request, templateGetter api.TemplateGetter, pathGetter interfaces.PathGetter) {
	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	day := vars["day"]
	name := vars["name"]
	paths, err := pathGetter.GetBlogPostPaths()
	if err != nil {
		log.Printf("error %v", err)
		api.InternalServerError(w, r)
		return
	}
	path := paths.GetPath(year, month, day, name)
	if len(path) == 0 {
		api.NotFound(w, r)
		return
	}
	blogPosts, err := blog.GetBlogPostData(path)
	if err != nil {
		log.Printf("error %v", err)
		api.InternalServerError(w, r)
		return
	}

	archiveLinks, err := blog.GetArchiveLinks(paths.GetPaths())
	if err != nil {
		log.Printf("error %v", err)
		api.InternalServerError(w, r)
		return
	}

	tmpl := templateGetter.GetTemplate()
	tmpl.Execute(w, api.BlogPage{
		TitleTag:       fmt.Sprintf("Charlton Austin's Blog Technically Dazed And Confused a blog post about: %v", name),
		DescriptionTag: fmt.Sprintf("Single article from: %v/%v/%v", year, month, day),
		HomeActive:     "active",
		AboutActive:    "",
		BlogPosts:      blogPosts,
		Previous:       paths.GetPrevious(),
		Next:           paths.GetNext(),
		ArchiveLinks:   archiveLinks,
	})

}
