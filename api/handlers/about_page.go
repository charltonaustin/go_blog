package handlers

import (
	"github.com/gomarkdown/markdown"
	"go-blog/api"
	"html/template"
	"net/http"
)

//AboutPage handler
type AboutPage struct {
	api.TemplateGetter
	api.AboutContentGetter
	api.IErrorHandler
}

//CreateAboutPageHandler returns a new about page handler
func CreateAboutPageHandler(
	templateGetter api.TemplateGetter,
	contentGetter api.AboutContentGetter,
	iErrorHandler api.IErrorHandler,
) AboutPage {
	return AboutPage{
		TemplateGetter:     templateGetter,
		AboutContentGetter: contentGetter,
		IErrorHandler:      iErrorHandler,
	}
}

func (a AboutPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content, err := a.GetContent()
	if err != nil {
		a.InternalServerError(w, r)
		return
	}
	output := markdown.ToHTML(content, nil, nil)
	tmpl := a.GetTemplate()
	err = tmpl.Execute(w, aboutPage{
		TitleTag:       "Charlton Austin's Blog Technically Dazed And Confused About Page",
		DescriptionTag: "A page describin who I am and what I'm doing with this blog.",
		Content:        template.HTML(output),
		HomeActive:     "",
		AboutActive:    "active",
	})
	if err != nil {
		a.InternalServerError(w, r)
	}
}

type aboutPage struct {
	DescriptionTag string
	TitleTag       string
	Content        template.HTML
	HomeActive     string
	AboutActive    string
}
