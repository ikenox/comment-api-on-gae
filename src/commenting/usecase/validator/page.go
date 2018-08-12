package validator

import (
	"errors"
	"regexp"
	"util"
)

var pageIdRegexp = regexp.MustCompile("^[0-9a-zA-Z_\\-]+$")

func ValidatePageID(pageId string) error {
	if pageId == "" {
		return errors.New("page ID must not be empty")
	}

	re := pageIdRegexp.Copy()
	pageId = re.FindString(pageId)
	if pageId == "" {
		return errors.New("invalid character ")
	}

	if len(pageId) > 64 {
		return errors.New("page ID is too long")
	}

	return nil
}

func ValidateComment(comment string) error {
	if comment == "" {
		return errors.New("comment must not be empty")
	}
	if util.LengthOf(comment) > 1000 {
		return errors.New("comment is too long")
	}
	return nil
}

func ValidateCommenterName(name string) error {
	if util.LengthOf(name) > 20 {
		return errors.New("commenter name is too long")
	}
	return nil
}
