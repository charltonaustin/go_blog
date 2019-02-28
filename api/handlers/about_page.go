package handlers

import (
	"github.com/gomarkdown/markdown"
	"go-api/api"
	"html/template"
	"net/http"
)

func CreateAboutPageHandler(templateGetter api.TemplateGetter, contentGetter api.ContentGetter) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AboutPage(w, r, templateGetter, contentGetter)
	})
}

func AboutPage(w http.ResponseWriter, r *http.Request, templateGetter api.TemplateGetter, contentGetter api.ContentGetter) {
	content, err := contentGetter.GetContent()
	if err != nil {
		api.InternalServerError(w, r)
		return
	}
	output := markdown.ToHTML(content, nil, nil)
	tmpl := templateGetter.GetTemplate()
	err = tmpl.Execute(w, aboutPage{
		TitleTag:       "Charlton Austin's Blog Technically Dazed And Confused About Page",
		DescriptionTag: "A page describin who I am and what I'm doing with this blog.",
		Content:        template.HTML(output),
		HomeActive:     "",
		AboutActive:    "active",
	})
	if err != nil {
		api.InternalServerError(w, r)
	}
}

type aboutPage struct {
	DescriptionTag string
	TitleTag       string
	Content        template.HTML
	HomeActive     string
	AboutActive    string
}
