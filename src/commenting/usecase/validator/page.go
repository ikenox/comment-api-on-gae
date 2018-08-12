package validator

import (
	"errors"
	"regexp"
	"util"
)

var pageIdRegexp = regexp.MustCompile("^[0-9a-zA-Z_\\-]+$")

func ValidatePageID(pageId string) error {
	if pageId == "" {
		return errors.New("PageID must not be empty")
	}

	re := pageIdRegexp.Copy()
	pageId = re.FindString(pageId)
	if pageId == "" {
		return errors.New("Invalid character ")
	}

	if len(pageId) > 64 {
		return errors.New("PageID is too long")
	}

	return nil
}

func ValidateComment(comment string) error {
	if comment == "" {
		return errors.New("Comment must not be empty")
	}
	if util.LengthOf(comment) > 1000 {
		return errors.New("Comment is too long")
	}
	return nil
}

func ValidateCommenterName(name string) error {
	if util.LengthOf(name) > 20 {
		return errors.New("Commenter name is too long")
	}
	return nil
}
