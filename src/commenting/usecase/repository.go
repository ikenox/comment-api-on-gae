package usecase

import (
	"commenting/domain"
)

type CommentRepository interface {
	NextCommentID() domain.CommentID
	Add(comment *domain.Comment)
	Delete(comment domain.CommentID)
	FindByPageID(page domain.PageID) []*domain.Comment
}

type PageRepository interface {
	Add(page *domain.Page)
	Delete(page domain.PageID)
	Get(pageId domain.PageID) *domain.Page
}

type CommenterRepository interface {
	NextCommenterID() domain.CommenterID
	FindByComments(commenterIDs []domain.CommenterID) []*domain.Commenter
	Add(commenter *domain.Commenter)
	Delete(commenterId domain.CommenterID)
	Get(commenterId domain.CommenterID) *domain.Commenter
}
