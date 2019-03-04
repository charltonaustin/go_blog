package blog

import (
	"go-blog/interfaces"
	"sort"
	"time"
)

//GetArchiveLinks returns archive links for the given post info
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

//NewArchiveLink creates new archive link
func NewArchiveLink(date time.Time) ArchiveLink {
	newDate := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, date.Location())
	return ArchiveLink{date: newDate}
}

//ArchiveLink returns link info for archived link
type ArchiveLink struct {
	date time.Time
}

//LinkDate returns formatted link for display
func (a ArchiveLink) LinkDate() string {
	return a.date.Format("January 2006")
}

//LinkHref returns link location for clicking
func (a ArchiveLink) LinkHref() string {
	return a.date.Format("/2006/01")
}
