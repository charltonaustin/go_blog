package interfaces

import (
	"html/template"
	"io"
	"time"
)

type Executor interface {
	Execute(wr io.Writer, data interface{}) error
}


type PathGetter interface {
	GetBlogPostPaths() (PostFinder, error)
}

type PostFinder interface {
	GetPaths() []PostInfo
	FromEnd(start int) []PostInfo
	GetPath(year, month, day, name string) []PostInfo
	GetArchive(year string, month string) []PostInfo
	GetNext() string
	GetPrevious() string
}
type PostInfo interface {
	Date() (*time.Time, error)
	Name()string
	Year() string
	Month() string
	Day() string
}

type IPost interface {
	Content() template.HTML
	Single() string
	GetDisplayName() string
	PublishDate() string
}

type IArchiveLink interface {
	LinkDate() string
	LinkHref() string
}