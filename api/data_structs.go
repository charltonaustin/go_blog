package api

import (
	"go-blog/interfaces"
)

//BlogPage data object
type BlogPage struct {
	DescriptionTag string
	Content        string
	TitleTag       string
	HomeActive     string
	AboutActive    string
	BlogPosts      []interfaces.IPost
	Previous       string
	Next           string
	ArchiveLinks   []interfaces.IArchiveLink
}
