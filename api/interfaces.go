package api

import "go-blog/interfaces"

type TemplateGetter interface {
	GetTemplate() interfaces.Executor
}

type ContentGetter interface {
	GetContent()([]byte, error)
}
