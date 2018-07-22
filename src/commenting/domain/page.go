package domain

import (
	"errors"
	"fmt"
	"regexp"
)

type PageId string

func NewPageId(pageId string) PageId {
	if err := PageIdSpec.CheckValidityOf(pageId); err != nil {
		panic(fmt.Sprintf("Invalid pageId: %s", err.Error()))
	}
	return PageId(pageId)
}

var pageIdRegexp = regexp.MustCompile("^[0-9a-zA-Z_\\-]+$")

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

type pageIdSpec struct{}

var PageIdSpec = &pageIdSpec{}

type PageIdValidityError error

var (
	ErrEmptyPageId      PageIdValidityError = errors.New("PageId is empty")
	ErrInvalidCharacter PageIdValidityError = errors.New("PageId contains invalid character")
	ErrPageIdTooLong    PageIdValidityError = errors.New("PageId is too long")
)

func (s *pageIdSpec) CheckValidityOf(pageId string) PageIdValidityError {
	if pageId == "" {
		return ErrEmptyPageId
	}

	re := pageIdRegexp.Copy()
	pageId = re.FindString(pageId)
	if pageId == "" {
		return ErrInvalidCharacter
	}

	if len(pageId) > 64 {
		return ErrPageIdTooLong
	}

	return nil
}
