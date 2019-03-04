package blog

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"go-blog/interfaces"
	"html/template"
	"io/ioutil"
	"strings"
	"time"
)

type ContentGetter struct {
	path string
}

func NewPostGetter(path string) ContentGetter {
	return ContentGetter{
		path: path,
	}
}

func (c ContentGetter) GetBlogPostData(paths []interfaces.PostInfo) ([]interfaces.IPost, error) {
	var blogPosts []interfaces.IPost
	single := len(paths) == 1
	for i := len(paths) - 1; i >= 0; i-- {
		blogPostInfo := paths[i]
		content, err := ioutil.ReadFile(fmt.Sprintf(
			c.path+"/published/%v/%v/%v/%v",
			blogPostInfo.Year(),
			blogPostInfo.Month(),
			blogPostInfo.Day(),
			blogPostInfo.Name()),
		)
		if err != nil {
			return nil, err
		}

		output := markdown.ToHTML(content, nil, nil)
		date, err := blogPostInfo.Date()
		if err != nil {
			return nil, err
		}
		blogPosts = append(blogPosts, NewBlogPost(single, blogPostInfo.Name(), *date, output))
	}
	return blogPosts, nil
}

func NewBlogPost(single bool, name string, publishDate time.Time, content []byte) interfaces.IPost {
	return interfaces.IPost(&Post{
		single:      single,
		name:        name,
		publishDate: publishDate,
		content:     content,
	})
}

type Post struct {
	single      bool
	name        string
	publishDate time.Time
	content     []byte
}

func (b Post) Content() template.HTML {
	return template.HTML(b.content)
}

func (b *Post) Single() string {
	if b.single {
		return "single-"
	}
	return ""
}

func (b *Post) GetDisplayName() string {
	name := strings.Replace(b.name, "-", " ", -1)
	name = strings.Replace(name, ".md", "", -1)
	name = strings.Title(name)
	return name
}

func (b *Post) PublishDate() string {
	return b.publishDate.Format("Jan 2, 2006")
}
