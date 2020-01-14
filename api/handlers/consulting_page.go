package handlers

import (
"github.com/gomarkdown/markdown"
"go-blog/api"
"html/template"
"net/http"
)

//ConsultingPage handler
type ConsultingPage struct {
	api.TemplateGetter
	api.AboutContentGetter
	api.IErrorHandler
}

//CreateConsultingPageHandler returns a new about page handler
func CreateConsultingPageHandler(
	templateGetter api.TemplateGetter,
	contentGetter api.AboutContentGetter,
	iErrorHandler api.IErrorHandler,
) ConsultingPage {
	return ConsultingPage{
		TemplateGetter:     templateGetter,
		AboutContentGetter: contentGetter,
		IErrorHandler:      iErrorHandler,
	}
}

func (a ConsultingPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content, err := a.GetContent()
	if err != nil {
		a.InternalServerError(w, r)
		return
	}
	output := markdown.ToHTML(content, nil, nil)
	tmpl := a.GetTemplate()
	err = tmpl.Execute(w, consultingPage{
		TitleTag:       "Charlton Austin's Blog Technically Dazed And Confused consulting Page",
		DescriptionTag: "A page describing the kind of consulting work I am available for.",
		Content:        template.HTML(output),
		HomeActive:     "",
		AboutActive:    "active",
	})
	if err != nil {
		a.InternalServerError(w, r)
	}
}

type consultingPage struct {
	DescriptionTag string
	TitleTag       string
	Content        template.HTML
	HomeActive     string
	AboutActive    string
}
