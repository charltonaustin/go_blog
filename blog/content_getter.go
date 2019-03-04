package blog

import "io/ioutil"

type AboutContentGetter struct {
	filePath string
}

func NewContentGetter(path string) AboutContentGetter {
	return AboutContentGetter{filePath: path}
}

func (c AboutContentGetter) GetContent() ([]byte, error) {
	return ioutil.ReadFile(c.filePath + "/about.md")
}
