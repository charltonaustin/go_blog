package handlers

import (
	"go-blog/api"
	"go-blog/blog"
	"net/http"
)

func CreateMainPageHandler(getter api.TemplateGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		MainPage(w, r, getter)
	})
}
func MainPage(w http.ResponseWriter, r *http.Request, getter api.TemplateGetter) {
	paths, err := blog.GetBlogPostPaths()
	if err != nil {
		api.InternalServerError(w, r)
		return
	}
	blogPosts, err := blog.GetBlogPostData(paths.FromEnd(3))
	if err != nil {
		api.InternalServerError(w, r)
		return
	}

	archiveLinks, err := blog.GetArchiveLinks(paths.GetPaths())
	if err != nil {
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
