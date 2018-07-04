package domain

import (
	"regexp"
)

type PageId string

var (
	InvalidPageIdError = Error{message: "invalid page id"}
)

func NewPageId(pageId string) PageId {
	if !IsValidPageId(pageId) {
		panic("Invalid pageId")
	}
	return PageId(pageId)
}

var pageIdRegexp = regexp.MustCompile("^[0-9a-zA-Z_\\-]+$")

func IsValidPageId(pageId string) bool {
	re := pageIdRegexp.Copy()
	pageId = re.FindString(pageId)
	return pageId != ""
}

type Page struct {
	pageId PageId
}

func NewPage(pageId PageId) *Page {
	return &Page{
		pageId: pageId,
	}
}

func (p *Page) PageId() PageId {
	return p.pageId
}
