package blog

import "io/ioutil"

//AboutContentGetter returns about content
type AboutContentGetter struct {
	filePath string
}

//NewContentGetter returns new
func NewContentGetter(path string) AboutContentGetter {
	return AboutContentGetter{filePath: path}
}

//GetContent returns the contents of the about page
func (c AboutContentGetter) GetContent() ([]byte, error) {
	return ioutil.ReadFile(c.filePath + "/about.md")
}
