package blog

import (
	"fmt"
	"go-blog/interfaces"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type PostGetter struct {
	path string
}

func NewBlogPostGetter(path string) PostGetter {
	return PostGetter{path}
}
func (pg PostGetter) GetBlogPostPaths() (interfaces.PostFinder, error) {
	pathMap := make(map[string]*PostPath)
	var paths []*PostPath
	fullPath := pg.path + "/blog-entries/published"
	err := filepath.Walk(
		fullPath,
		filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && strings.Contains(info.Name(), ".md") {
				path = strings.Replace(
					path,
					fullPath + "/",
					"",
					-1,
				)
				date := strings.Split(path, "/")
				postPath := PostPath{
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
	pathMap  map[string]*PostPath
	paths    []*PostPath
	next     int
	previous int
}

func (p *PostPaths) GetPaths() []interfaces.PostInfo {
	var postInfo []interfaces.PostInfo
	for _, p := range p.paths {
		postInfo = append(postInfo, interfaces.PostInfo(p))
	}
	return postInfo
}

func (p *PostPaths) FromEnd(start int) []interfaces.PostInfo {
	index := len(p.paths) - start
	if index < 0 {
		index = 0
	}

	paths := p.paths[index:]
	p.next = p.paths[index].location - 1
	p.previous = paths[len(paths)-1].location + 1
	return postInfo(p)
}

func (p *PostPaths) GetPath(year, month, day, name string) []interfaces.PostInfo {
	path := p.pathMap[fmt.Sprintf("%v:%v:%v:%v", year, month, day, name)]
	if path == nil {
		return []interfaces.PostInfo{}
	}

	p.next = path.location - 1
	p.previous = path.location + 1
	return postInfo(p)
}

func postInfo(p *PostPaths) []interfaces.PostInfo {
	var postInfo []interfaces.PostInfo
	for _, p := range p.paths {
		postInfo = append(postInfo, interfaces.PostInfo(p))
	}
	return postInfo
}

func (p *PostPaths) GetArchive(year string, month string) []interfaces.PostInfo {
	var postPaths []*PostPath
	for key, value := range p.pathMap {
		if strings.Contains(key, fmt.Sprintf("%v:%v", year, month)) {
			postPaths = append(postPaths, value)
		}
	}
	if postPaths == nil {
		return []interfaces.PostInfo{}
	}

	path := postPaths[len(postPaths)-1]
	next := path.location - 1
	path = postPaths[0]
	previous := path.location + 1
	p.next = next
	p.previous = previous
	return postInfo(p)
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

type PostPath struct {
	year     string
	month    string
	day      string
	name     string
	location int
}

func (b PostPath) Date() (*time.Time, error) {
	year, err := strconv.Atoi(b.year)
	if err != nil {
		log.Printf("failed to parse year %v with Atoi", b.Year())
		return nil, err
	}

	month, err := strconv.Atoi(b.month)
	if err != nil {
		log.Printf("failed to parse month %v with Atoi", b.Month())
		return nil, err
	}

	day, err := strconv.Atoi(b.day)
	if err != nil {
		log.Printf("failed to parse day %v with Atoi", b.Day())
		return nil, err
	}
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location())
	return &date, nil
}

func (b PostPath) Name() string {
	return b.name
}

func (b PostPath) Year() string {
	return b.year
}

func (b PostPath) Month() string {
	return b.month
}


func (b PostPath) Day() string {
	return b.day
}