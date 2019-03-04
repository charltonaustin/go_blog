package api

import (
	"go-blog/interfaces"
	"net/http"
)

//TemplateGetter interface
type TemplateGetter interface {
	GetTemplate() interfaces.Executor
}

//AboutContentGetter interface
type AboutContentGetter interface {
	GetContent() ([]byte, error)
}

//IErrorHandler interface
type IErrorHandler interface {
	NotFound(w http.ResponseWriter, r *http.Request)
	InternalServerError(w http.ResponseWriter, r *http.Request)
}
