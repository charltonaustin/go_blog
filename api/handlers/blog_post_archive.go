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

//BlogPostArchive archived blog post handler
type BlogPostArchive struct {
	api.TemplateGetter
	interfaces.PathGetter
	api.IErrorHandler
	contentGetter
}

//CreateBlogPostArchiveHandler returns new BlogPostArchive
func CreateBlogPostArchiveHandler(
	getter api.TemplateGetter,
	pathGetter interfaces.PathGetter,
	errorHandler api.IErrorHandler,
	contentGetter contentGetter,
) BlogPostArchive {
	return BlogPostArchive{
		TemplateGetter: getter,
		PathGetter:     pathGetter,
		IErrorHandler:  errorHandler,
		contentGetter:  contentGetter,
	}
}

func (b BlogPostArchive) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths, err := b.GetBlogPostPaths()
	if err != nil {
		log.Printf("error %v", err)
		b.InternalServerError(w, r)
		return
	}

	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	archive := paths.GetArchive(year, month)
	if len(archive) == 0 {
		b.NotFound(w, r)
		return
	}

	posts, err := b.GetBlogPostData(archive)
	if err != nil {
		log.Printf("error %v", err)
		b.InternalServerError(w, r)
		return
	}

	archiveLinks, err := blog.GetArchiveLinks(paths.GetPaths())
	if err != nil {
		log.Printf("error %v", err)
		b.InternalServerError(w, r)
		return
	}

	tmpl := b.GetTemplate()
	title := fmt.Sprintf(
		"Charlton Austin's Blog Technically Dazed And Confused blog posts from: %v/%v",
		year,
		month,
	)
	tmpl.Execute(w, api.BlogPage{
		TitleTag:       title,
		DescriptionTag: fmt.Sprintf("This is all the blog posts I wrote from %v/%v", year, month),
		HomeActive:     "active",
		AboutActive:    "",
		BlogPosts:      posts,
		Previous:       paths.GetPrevious(),
		Next:           paths.GetNext(),
		ArchiveLinks:   archiveLinks,
	})
}
