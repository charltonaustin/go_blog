package blog

import "io/ioutil"

type ContentGetter struct {

}

func (c ContentGetter)GetContent()([]byte, error){
	return ioutil.ReadFile("/Users/charltonaustin/dev/personal/blog-entries/about.md")
}
