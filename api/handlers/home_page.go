package handlers

import (
	"go-blog/api"
	"go-blog/blog"
	"go-blog/interfaces"
	"log"
	"net/http"
)

type MainPage struct {
	api.TemplateGetter
	interfaces.PathGetter
	api.IErrorHandler
	contentGetter
}

func CreateMainPageHandler(
	getter api.TemplateGetter,
	pathGetter interfaces.PathGetter,
	errorHandler api.IErrorHandler,
	contentGetter contentGetter,
) MainPage {
	return MainPage{
		TemplateGetter: getter,
		PathGetter:     pathGetter,
		IErrorHandler:  errorHandler,
		contentGetter: contentGetter,
	}
}
func (m MainPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths, err := m.GetBlogPostPaths()
	if err != nil {
		log.Printf("error %v", err)
		m.InternalServerError(w, r)
		return
	}
	blogPosts, err := m.GetBlogPostData(paths.FromEnd(3))
	if err != nil {
		log.Printf("error %v", err)
		m.InternalServerError(w, r)
		return
	}

	archiveLinks, err := blog.GetArchiveLinks(paths.GetPaths())
	if err != nil {
		log.Printf("error %v", err)
		m.InternalServerError(w, r)
		return
	}

	tmpl := m.GetTemplate()
	tmpl.Execute(w, api.BlogPage{
		TitleTag:       "Charlton Austin's Blog Technically Dazed And Confused Home Page",
		DescriptionTag: "The landing page for my blog. It contains the three latest blog posts that I have written.",
		HomeActive:     "active",
		AboutActive:    "",
		BlogPosts:      blogPosts,
		Previous:       paths.GetPrevious(),
		Next:           paths.GetNext(),
		ArchiveLinks:   archiveLinks,
	})
}
