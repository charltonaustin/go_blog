package handlers

import (
	"go-blog/interfaces"
)

type contentGetter interface {
	GetBlogPostData(paths []interfaces.PostInfo) ([]interfaces.IPost, error)
}
