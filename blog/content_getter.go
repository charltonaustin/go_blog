package blog

import "io/ioutil"

type ContentGetter struct {
	filePath string
}

func NewContentGetter(path string) ContentGetter{
	return ContentGetter{filePath: path}
}

func (c ContentGetter) GetContent() ([]byte, error) {
	return ioutil.ReadFile(c.filePath + "/blog-entries/about.md")
}
