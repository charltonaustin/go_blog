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

func CreateBlogPostArchiveHandler(getter api.TemplateGetter, pathGetter interfaces.PathGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		BlogPostArchive(w, r, getter, pathGetter)
	})
}
func BlogPostArchive(w http.ResponseWriter, r *http.Request, getter api.TemplateGetter, pathGetter interfaces.PathGetter) {
	paths, err := pathGetter.GetBlogPostPaths()
	if err != nil {
		log.Printf("error %v", err)
		api.InternalServerError(w, r)
		return
	}

	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	archive := paths.GetArchive(year, month)
	if len(archive) == 0 {
		api.NotFound(w, r)
		return
	}

	posts, err := blog.GetBlogPostData(archive)
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

	tmpl := getter.GetTemplate()
	tmpl.Execute(w, api.BlogPage{
		TitleTag:       fmt.Sprintf("Charlton Austin's Blog Technically Dazed And Confused blog posts from: %v/%v", year, month),
		DescriptionTag: fmt.Sprintf("This is all the blog posts I wrote from %v/%v", year, month),
		HomeActive:     "active",
		AboutActive:    "",
		BlogPosts:      posts,
		Previous:       paths.GetPrevious(),
		Next:           paths.GetNext(),
		ArchiveLinks:   archiveLinks,
	})
}
