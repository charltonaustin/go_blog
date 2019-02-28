package api

import (
	"go-api/blog"
)

type BlogPage struct {
	DescriptionTag string
	Content        string
	TitleTag       string
	HomeActive     string
	AboutActive    string
	BlogPosts      []*blog.Post
	Previous       string
	Next           string
	ArchiveLinks   []blog.ArchiveLink
}

