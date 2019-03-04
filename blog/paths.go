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

//PostGetter object for getting post data
type PostGetter struct {
	path string
}

//NewPathGetter returns new instance
func NewPathGetter(path string) PostGetter {
	return PostGetter{path}
}

//GetBlogPostPaths returns the blog post paths
func (pg PostGetter) GetBlogPostPaths() (interfaces.PostFinder, error) {
	pathMap := make(map[string]*postPath)
	var paths []*postPath
	fullPath := pg.path + "/published"
	err := filepath.Walk(
		fullPath,
		filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && strings.Contains(info.Name(), ".md") {
				path = strings.Replace(
					path,
					fullPath+"/",
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
	return &postPaths{pathMap: pathMap, paths: paths}, nil
}

type postPaths struct {
	pathMap  map[string]*postPath
	paths    []*postPath
	next     int
	previous int
}

func (p *postPaths) GetPaths() []interfaces.PostInfo {
	var postInfo []interfaces.PostInfo
	for _, p := range p.paths {
		postInfo = append(postInfo, interfaces.PostInfo(p))
	}
	return postInfo
}

func (p *postPaths) FromEnd(start int) []interfaces.PostInfo {
	index := len(p.paths) - start
	if index < 0 {
		index = 0
	}

	paths := p.paths[index:]
	p.next = p.paths[index].location - 1
	p.previous = paths[len(paths)-1].location + 1
	return postInfo(paths)
}

func (p *postPaths) GetPath(year, month, day, name string) []interfaces.PostInfo {
	path := p.pathMap[fmt.Sprintf("%v:%v:%v:%v", year, month, day, name)]
	if path == nil {
		return []interfaces.PostInfo{}
	}

	p.next = path.location - 1
	p.previous = path.location + 1
	return postInfo([]*postPath{path})
}

func postInfo(paths []*postPath) []interfaces.PostInfo {
	var postInfo []interfaces.PostInfo
	for _, p := range paths {
		postInfo = append(postInfo, interfaces.PostInfo(p))
	}
	return postInfo
}

func (p *postPaths) GetArchive(year string, month string) []interfaces.PostInfo {
	var postPaths []*postPath
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
	return postInfo(postPaths)
}

func (p *postPaths) GetNext() string {
	index := p.next
	if index < 0 {
		index = 0
	}
	path := p.paths[index]
	return fmt.Sprintf("/%v/%v/%v/%v", path.year, path.month, path.day, path.name)
}

func (p *postPaths) GetPrevious() string {
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

func (b postPath) Name() string {
	return b.name
}

func (b postPath) Year() string {
	return b.year
}

func (b postPath) Month() string {
	return b.month
}

func (b postPath) Day() string {
	return b.day
}
