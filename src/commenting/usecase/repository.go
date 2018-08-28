package usecase

import (
	"commenting/domain/auth"
	"commenting/domain/comment"
)

type CommentRepository interface {
	NextCommentID() comment.CommentID
	Add(comment *comment.Comment)
	Delete(comment comment.CommentID)
	FindByPageID(page comment.PageID) []*comment.Comment
}

type PageRepository interface {
	Add(page *comment.Page)
	Delete(page comment.PageID)
	Get(pageId comment.PageID) *comment.Page
}

type CommenterRepository interface {
	NextCommenterID() comment.CommenterID
	FindByComments(commenterIDs []comment.CommenterID) []*comment.Commenter
	Add(commenter *comment.Commenter)
	Delete(commenterId comment.CommenterID)
	Get(commenterId comment.CommenterID) *comment.Commenter
}

type UserRepository interface {
	CurrentUser() *auth.User
}
