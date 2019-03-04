package interfaces

import (
	"html/template"
	"io"
	"time"
)

//Executor execute or error
type Executor interface {
	Execute(wr io.Writer, data interface{}) error
}

//PathGetter get paths or error
type PathGetter interface {
	GetBlogPostPaths() (PostFinder, error)
}

//PostFinder all methods needed for a post
type PostFinder interface {
	GetPaths() []PostInfo
	FromEnd(start int) []PostInfo
	GetPath(year, month, day, name string) []PostInfo
	GetArchive(year string, month string) []PostInfo
	GetNext() string
	GetPrevious() string
}

//PostInfo methods for post info
type PostInfo interface {
	Date() (*time.Time, error)
	Name() string
	Year() string
	Month() string
	Day() string
}

//IPost methods for post data
type IPost interface {
	Content() template.HTML
	Single() string
	GetDisplayName() string
	PublishDate() string
}

//IArchiveLink methods for archived links
type IArchiveLink interface {
	LinkDate() string
	LinkHref() string
}
