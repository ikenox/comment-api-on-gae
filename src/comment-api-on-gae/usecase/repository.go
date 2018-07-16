package usecase

import "comment-api-on-gae/domain"

type CommentRepository interface {
	NextCommentId() domain.CommentId
	Add(comment *domain.Comment)
	Delete(comment domain.CommentId)
	FindByPageId(page domain.PageId) []*domain.Comment
}

type PageRepository interface {
	NextPageId() domain.PageId
	Add(page *domain.Page)
	Delete(page domain.PageId)
	Get(pageId domain.PageId) *domain.Page
}

type CommenterRepository interface {
	NextCommenterId() domain.CommenterId
	FindByComments(commenterIds []domain.CommenterId) []*domain.Commenter
	Add(commenter *domain.Commenter)
	Delete(commenterId domain.CommenterId)
	Get(commenterId domain.CommenterId) *domain.Commenter
}
