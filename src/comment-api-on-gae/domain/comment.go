package domain

import "time"

type CommentId int
type Comment struct {
	Entity
	pageId      *PageId
	text        string
	commenterId *CommenterId
	commentedAt time.Time
	isDeleted   bool
}