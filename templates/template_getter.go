package templates

import (
	"go-blog/interfaces"
	"html/template"
)

//ErrorTemplateGetter gets templates for error handler
type ErrorTemplateGetter struct {
	path string
}

//NewErrorTemplateGetter returns new
func NewErrorTemplateGetter(path string) ErrorTemplateGetter {
	return ErrorTemplateGetter{path: path}
}

//GetTemplate for errors
func (b ErrorTemplateGetter) GetTemplate() interfaces.Executor {
	return template.Must(template.ParseFiles(b.path + "/templates/error.html"))
}

//BlogTemplateGetter gets templates for blog post
type BlogTemplateGetter struct {
	path string
}

//NewBlogTemplateGetter returns new
func NewBlogTemplateGetter(path string) BlogTemplateGetter {
	return BlogTemplateGetter{path: path}
}

//GetTemplate gets templates for blog post
func (b BlogTemplateGetter) GetTemplate() interfaces.Executor {
	return template.Must(template.ParseFiles(
		b.path+"/templates/base.html",
		b.path+"/templates/home.html",
		b.path+"/templates/blog.html",
		b.path+"/templates/sidebar.html",
		b.path+"/templates/pager.html",
	))
}

//AboutPageGetter gets templates for about page
type AboutPageGetter struct {
	path string
}

//NewAboutPageGetter returns new
func NewAboutPageGetter(path string) AboutPageGetter {
	return AboutPageGetter{path: path}
}

//GetTemplate gets templates for about page
func (a AboutPageGetter) GetTemplate() interfaces.Executor {
	return template.Must(template.ParseFiles(
		a.path+"/templates/base.html",
		a.path+"/templates/about.html",
	))
}

//ConsultingPageGetter gets templates for about page
type ConsultingPageGetter struct {
	path string
}

//NewAboutPageGetter returns new
func NewConsultingPageGetter(path string) ConsultingPageGetter {
	return ConsultingPageGetter{path: path}
}

//GetTemplate gets templates for about page
func (a ConsultingPageGetter) GetTemplate() interfaces.Executor {
	return template.Must(template.ParseFiles(
		a.path+"/templates/base.html",
		a.path+"/templates/consulting.html",
	))
}