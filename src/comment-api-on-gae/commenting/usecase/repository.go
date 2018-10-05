package usecase

import (
	"comment-api-on-gae/commenting/domain"
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
	NextCommenterID() domain.CommenterID
	FindByCommenterID(userIDs []domain.CommenterID) []*domain.Commenter
	CurrentUser(idToken string) domain.UserID
	Put(commenter *domain.Commenter)
}

type EventPublisher interface {
	Publish(topicID string, message string)
}

type LoggingRepository interface {
	Infof(format string, args ...interface{})
}
