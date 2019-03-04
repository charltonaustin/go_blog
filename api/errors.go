package api

import (
	"html/template"
	"net/http"
)

func NotFoundHandlerCreator() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		NotFound(w, r)
	})

}

func NotFound(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	tmpl.Execute(w, struct {
		Title   string
		Message string
		Status  string
	}{
		Title:   "Page not found",
		Message: "We could not find the page you were looking for",
		Status:  "404",
	})
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/error.html"))
	tmpl.Execute(w, struct {
		Title   string
		Message string
		Status  string
	}{
		Title:   "Uh, oh!",
		Message: "Something went terribly wrong. Please try again later. If this problem persists please contact me.",
		Status:  "500",
	})
}