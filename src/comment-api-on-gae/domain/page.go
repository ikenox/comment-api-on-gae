package domain

import (
	"regexp"
)

type PageId string

type InvalidPageIdError Error

func NewPageId(pageId string) (PageId, *Error) {
	if !isValidPageId(pageId) {
		err := Error(InvalidPageIdError{message: "invalid PageId"})
		return "", &err
	}
	return PageId(pageId), nil
}

var pageIdRegexp = regexp.MustCompile("^[0-9a-zA-Z_\\-]+$")

func isValidPageId(pageId string) bool {
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
