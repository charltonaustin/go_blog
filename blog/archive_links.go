package blog

import (
	"go-blog/interfaces"
	"sort"
	"time"
)

func GetArchiveLinks(paths []interfaces.PostInfo) ([]interfaces.IArchiveLink, error) {
	archiveSet := make(map[string]ArchiveLink)
	for _, bpp := range paths {
		date, err := bpp.Date()
		if err != nil {
			return nil, err
		}
		archiveSet[date.Format("2006/01")] = NewArchiveLink(*date)
	}
	var archiveLinks []ArchiveLink
	for _, archiveLink := range archiveSet {
		archiveLinks = append(archiveLinks, archiveLink)
	}
	sort.Slice(archiveLinks, func(i, j int) bool {
		return archiveLinks[i].date.Before(archiveLinks[j].date)
	})

	var links []interfaces.IArchiveLink
	for _, link := range archiveLinks {
		links = append(links, interfaces.IArchiveLink(link))
	}
	return links, nil
}

func NewArchiveLink(date time.Time) ArchiveLink {
	newDate := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, date.Location())
	return ArchiveLink{date: newDate}
}

type ArchiveLink struct {
	date time.Time
}

func (a ArchiveLink) LinkDate() string {
	return a.date.Format("January 2006")
}

func (a ArchiveLink) LinkHref() string {
	return a.date.Format("/2006/01")
}
