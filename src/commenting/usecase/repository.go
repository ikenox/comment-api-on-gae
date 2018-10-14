package usecase

import (
	"commenting/domain"
)

type CommentRepository interface {
	NextCommentID() domain.CommentID
	Get(pageId domain.CommentID) *domain.Comment
	Put(comment *domain.Comment)
	Delete(commentID domain.CommentID)
	FindByPageID(page domain.PageID) []*domain.Comment
}

type PageRepository interface {
	Add(page *domain.Page)
	Delete(page domain.PageID)
	Get(pageId domain.PageID) *domain.Page
}

type CommenterRepository interface {
	CurrentCommenter(idToken string) *domain.Commenter
}

type EventPublisher interface {
	Publish(event string, src interface{})
}

type LoggingRepository interface {
	Infof(format string, args ...interface{})
}
