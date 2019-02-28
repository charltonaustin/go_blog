package api

import "go-api/interfaces"

type TemplateGetter interface {
	GetTemplate() interfaces.Executor
}

type ContentGetter interface {
	GetContent()([]byte, error)
}
