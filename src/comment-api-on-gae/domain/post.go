package domain

import (
	"time"
)

type Page struct {
	url string
}

// Comment stores posted comment
type Post struct {
	Entity
	pageId    int64
	text      string
	posterId  int64
	isDeleted bool
	postedAt  time.Time
}
