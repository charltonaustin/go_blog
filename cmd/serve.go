package main

import (
	"github.com/gorilla/mux"
	"go-blog/api"
	"go-blog/api/handlers"
	"go-blog/api/middleware"
	"go-blog/blog"
	"go-blog/templates"
	"log"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	blogPostPath := getEnv("BLOG_POSTS", "/Users/charltonaustin/dev/personal")
	router := mux.NewRouter()
	router.
		Methods("GET").
		Path("/{year}/{month}/{day}/{name}").
		Handler(handlers.CreateSpecificBlogPostHandler(
			templates.BlogTemplateGetter{},
			blog.NewBlogPostGetter(blogPostPath),
		))

	router.
		Methods("GET").
		Path("/{year}/{month}").
		Handler(handlers.CreateBlogPostArchiveHandler(
			templates.BlogTemplateGetter{},
			blog.NewBlogPostGetter(blogPostPath),
		))

	router.
		Methods("GET").
		Path("/").
		Handler(handlers.CreateMainPageHandler(templates.BlogTemplateGetter{}, blog.NewBlogPostGetter(blogPostPath)))

	getter := blog.NewContentGetter(blogPostPath)
	router.
		Methods("GET").
		Path("/about").
		Handler(handlers.CreateAboutPageHandler(templates.AboutPageGetter{}, getter))

	router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/public"))))

	router.
		NotFoundHandler = api.NotFoundHandlerCreator()

	router.Use(middleware.AccessLogger, middleware.RecoverFromPanic)
	log.Printf("Listening on %v; ctrl + c to stop", ":9000")
	http.ListenAndServe(":9000", router)
}
