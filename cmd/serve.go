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
	blogPostPath := getEnv("BLOG_POSTS", "/Users/charltonaustin/dev/personal/blog-entries")

	// Set up dependencies
	postGetter := blog.NewPostGetter(blogPostPath)
	errorHandler := api.NewErrorHandler(templates.NewErrorTemplateGetter(blogPostPath))
	pathGetter := blog.NewPathGetter(blogPostPath)
	templateGetter := templates.NewBlogTemplateGetter(blogPostPath)

	// Set up routes
	router := mux.NewRouter()
	router.
		Methods("GET").
		Path("/{year}/{month}/{day}/{name}").
		Handler(handlers.CreateSpecificBlogPostHandler(
			templateGetter,
			pathGetter,
			errorHandler,
			postGetter,
		))

	router.
		Methods("GET").
		Path("/{year}/{month}").
		Handler(handlers.CreateBlogPostArchiveHandler(
			templateGetter,
			pathGetter,
			errorHandler,
			postGetter,
		))

	router.
		Methods("GET").
		Path("/").
		Handler(handlers.CreateMainPageHandler(
			templateGetter,
			pathGetter,
			errorHandler,
			postGetter,
		))

	router.
		Methods("GET").
		Path("/about").
		Handler(handlers.CreateAboutPageHandler(
			templates.NewAboutPageGetter(blogPostPath),
			blog.NewContentGetter(blogPostPath),
			errorHandler,
		))

	router.
		Methods("GET").
		Path("/consulting").
		Handler(handlers.CreateConsultingPageHandler(
			templates.NewConsultingPageGetter(blogPostPath),
			blog.NewContentGetter(blogPostPath),
			errorHandler,
		))

	router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(blogPostPath+"static/public"))))

	router.NotFoundHandler = http.HandlerFunc(errorHandler.NotFound)

	router.Use(
		middleware.AccessLogger,
		middleware.NewRecoverHandler(errorHandler).RecoverFromPanic,
	)

	port := getEnv("BLOG_PORT", ":3000")
	log.Printf("Listening on %v; ctrl + c to stop", port)
	http.ListenAndServe(port, router)
}
