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

type SpecificBlogPost struct {
	api.TemplateGetter
	interfaces.PathGetter
	api.IErrorHandler
	contentGetter

}

func CreateSpecificBlogPostHandler(
	templateGetter api.TemplateGetter,
	pathGetter interfaces.PathGetter,
	handler api.ErrorHandler,
	contentGetter contentGetter,
) SpecificBlogPost {
	return SpecificBlogPost{
		TemplateGetter: templateGetter,
		PathGetter:     pathGetter,
		IErrorHandler:  handler,
		contentGetter: contentGetter,
	}
}
func (s SpecificBlogPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	day := vars["day"]
	name := vars["name"]
	paths, err := s.GetBlogPostPaths()
	if err != nil {
		log.Printf("error %v", err)
		s.InternalServerError(w, r)
		return
	}
	path := paths.GetPath(year, month, day, name)
	if len(path) == 0 {
		s.NotFound(w, r)
		return
	}
	blogPosts, err := s.GetBlogPostData(path)
	if err != nil {
		log.Printf("error %v", err)
		s.InternalServerError(w, r)
		return
	}

	archiveLinks, err := blog.GetArchiveLinks(paths.GetPaths())
	if err != nil {
		log.Printf("error %v", err)
		s.InternalServerError(w, r)
		return
	}

	tmpl := s.GetTemplate()
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