package templates

import (
	"go-blog/interfaces"
	"html/template"
)

type ErrorTemplateGetter struct {
}

type BlogTemplateGetter struct {
}

func (b BlogTemplateGetter) GetTemplate() interfaces.Executor {
	return template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/home.html",
		"templates/blog.html",
		"templates/sidebar.html",
		"templates/pager.html",
	))
}

type AboutPageGetter struct {
}

func (a AboutPageGetter) GetTemplate() interfaces.Executor {
	return template.Must(template.ParseFiles("templates/base.html", "templates/about.html"))
}
