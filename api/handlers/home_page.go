package handlers

import (
	"go-blog/api"
	"go-blog/blog"
	"go-blog/interfaces"
	"log"
	"net/http"
)

func CreateMainPageHandler(getter api.TemplateGetter, pathGetter interfaces.PathGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		MainPage(w, r, getter, pathGetter)
	})
}
func MainPage(w http.ResponseWriter, r *http.Request, getter api.TemplateGetter, pathGetter interfaces.PathGetter) {
	paths, err := pathGetter.GetBlogPostPaths()
	if err != nil {
		log.Printf("error %v", err)
		api.InternalServerError(w, r)
		return
	}
	blogPosts, err := blog.GetBlogPostData(paths.FromEnd(3))
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
