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
)

func main() {
	router := mux.NewRouter()
	router.
		Methods("GET").
		Path("/{year}/{month}/{day}/{name}").
		Handler(handlers.CreateSpecificBlogPostHandler(templates.BlogTemplateGetter{}))

	router.
		Methods("GET").
		Path("/{year}/{month}").
		Handler(handlers.CreateBlogPostArchiveHandler(templates.BlogTemplateGetter{}))

	router.
		Methods("GET").
		Path("/").
		Handler(handlers.CreateMainPageHandler(templates.BlogTemplateGetter{}))

	router.
		Methods("GET").
		Path("/about").
		Handler(handlers.CreateAboutPageHandler(templates.AboutPageGetter{}, blog.ContentGetter{}))

	router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/public"))))

	router.
		NotFoundHandler = api.NotFoundHandlerCreator()

	router.Use(middleware.AccessLogger, middleware.RecoverFromPanic)
	log.Printf("Listening on %v; ctrl + c to stop", ":9000")
	http.ListenAndServe(":9000", router)
}
