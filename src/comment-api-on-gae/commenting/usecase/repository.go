package usecase

import (
	"comment-api-on-gae/commenting/domain"
)

type CommentRepository interface {
	NextCommentID() domain.CommentID
	Put(comment *domain.Comment)
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
	FindByCommenterID(userIDs []domain.CommenterID) []*domain.Commenter
	CurrentCommenter(idToken string) *domain.Commenter
	Put(commenter *domain.Commenter)
}
