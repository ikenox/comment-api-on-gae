package domain

import (
	"time"
)

// Comment stores posted comment
type Post struct {
	Entity
	text      string
	posterId  int64
	isDeleted bool
	postedAt  time.Time
}
