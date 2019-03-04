package templates

import (
	"go-blog/interfaces"
	"html/template"
)

type ErrorTemplateGetter struct {
}

type BlogTemplateGetter struct {
	path string
}

func NewBlogTemplateGetter(path string) BlogTemplateGetter {
	return BlogTemplateGetter{path: path}
}


func (b BlogTemplateGetter) GetTemplate() interfaces.Executor {
	return template.Must(template.ParseFiles(
		b.path+"/templates/base.html",
		b.path+"/templates/home.html",
		b.path+"/templates/blog.html",
		b.path+"/templates/sidebar.html",
		b.path+"/templates/pager.html",
	))
}

type AboutPageGetter struct {
	path string
}

func NewAboutPageGetter(path string) AboutPageGetter {
	return AboutPageGetter{path: path}
}
func (a AboutPageGetter) GetTemplate() interfaces.Executor {
	return template.Must(template.ParseFiles(
		a.path + "/templates/base.html",
		a.path + "/templates/about.html",
	))
}
