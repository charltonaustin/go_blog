package api

import (
	"net/http"
)

//ErrorHandler deals with api level errors
type ErrorHandler struct {
	TemplateGetter
}

//NewErrorHandler returns a new error handler
func NewErrorHandler(templateGetter TemplateGetter) ErrorHandler {
	return ErrorHandler{TemplateGetter: templateGetter}
}

//ErrorInfo data struct for error info
type ErrorInfo struct {
	Title   string
	Message string
	Status  string
}

//NotFound returns template for 404
func (e ErrorHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	tmpl := e.GetTemplate()
	tmpl.Execute(w, ErrorInfo{
		Title:   "Page not found",
		Message: "We could not find the page you were looking for",
		Status:  "404",
	})
}

//InternalServerError returns template for 500
func (e ErrorHandler) InternalServerError(w http.ResponseWriter, r *http.Request) {
	tmpl := e.GetTemplate()
	tmpl.Execute(w, ErrorInfo{
		Title:   "Uh, oh!",
		Message: "Something went terribly wrong. Please try again later. If this problem persists please contact me.",
		Status:  "500",
	})
}
