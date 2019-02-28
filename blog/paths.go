package blog

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)



func GetBlogPostPaths() (*PostPaths, error) {
	pathMap := make(map[string]*postPath)
	var paths []*postPath
	err := filepath.Walk(
		"/Users/charltonaustin/dev/personal/blog-entries/published",
		filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && strings.Contains(info.Name(), ".md") {
				path = strings.Replace(
					path,
					"/Users/charltonaustin/dev/personal/blog-entries/published/",
					"",
					-1,
				)
				date := strings.Split(path, "/")
				postPath := postPath{
					year:  date[0],
					month: date[1],
					day:   date[2],
					name:  info.Name(),
				}
				_, err := postPath.Date()
				if err != nil {
					return err
				}
				pathMap[fmt.Sprintf("%v:%v:%v:%v", date[0], date[1], date[2], info.Name())] = &postPath
				paths = append(paths, &postPath)
			}
			return nil
		}))
	if err != nil {
		return nil, err
	}
	sort.Slice(paths, func(i, j int) bool {
		dateI, _ := paths[i].Date()
		dateJ, _ := paths[j].Date()
		return dateI.Before(*dateJ)
	})
	for i, p := range paths {
		key := fmt.Sprintf("%v:%v:%v:%v", p.year, p.month, p.day, p.name)
		pathMap[key].location = i
		p.location = i
	}
	return &PostPaths{pathMap: pathMap, paths: paths}, nil
}

type PostPaths struct {
	pathMap  map[string]*postPath
	paths    []*postPath
	next     int
	previous int
}

func (p *PostPaths) GetPaths() []*postPath {
	return p.paths
}

func (p *PostPaths) FromEnd(start int) []*postPath {
	index := len(p.paths) - start
	if index < 0 {
		index = 0
	}

	paths := p.paths[index:]
	p.next = p.paths[index].location - 1
	p.previous = paths[len(paths)-1].location + 1
	return paths
}

func (p *PostPaths) GetPath(year, month, day, name string) []*postPath {
	path := p.pathMap[fmt.Sprintf("%v:%v:%v:%v", year, month, day, name)]
	if path == nil {
		return []*postPath{}
	}

	p.next = path.location - 1
	p.previous = path.location + 1
	return []*postPath{path}
}

func (p *PostPaths) GetArchive(year string, month string) []*postPath {
	var postPaths []*postPath
	for key, value := range p.pathMap {
		if strings.Contains(key, fmt.Sprintf("%v:%v", year, month)) {
			postPaths = append(postPaths, value)
		}
	}
	if postPaths == nil {
		return []*postPath{}
	}

	path := postPaths[len(postPaths)-1]
	next := path.location - 1
	path = postPaths[0]
	previous := path.location + 1
	p.next = next
	p.previous = previous
	return postPaths
}

func (p *PostPaths) GetNext() string {
	index := p.next
	if index < 0 {
		index = 0
	}
	path := p.paths[index]
	return fmt.Sprintf("/%v/%v/%v/%v", path.year, path.month, path.day, path.name)
}

func (p *PostPaths) GetPrevious() string {
	if p.previous >= len(p.paths) {
		return "/"
	}
	path := p.paths[p.previous]
	return fmt.Sprintf("/%v/%v/%v/%v", path.year, path.month, path.day, path.name)

}

type postPath struct {
	year     string
	month    string
	day      string
	name     string
	location int
}

func (b postPath) Date() (*time.Time, error) {
	year, err := strconv.Atoi(b.year)
	if err != nil {
		return nil, err
	}

	month, err := strconv.Atoi(b.month)
	if err != nil {
		return nil, err
	}

	day, err := strconv.Atoi(b.day)
	if err != nil {
		return nil, err
	}
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location())
	return &date, nil
}
