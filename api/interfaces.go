package api

import (
	"go-blog/interfaces"
	"net/http"
)

type TemplateGetter interface {
	GetTemplate() interfaces.Executor
}

type AboutContentGetter interface {
	GetContent() ([]byte, error)
}

type IErrorHandler interface {
	NotFound(w http.ResponseWriter, r *http.Request)
	InternalServerError(w http.ResponseWriter, r *http.Request)
}