package usecase

import (
	"comment-api-on-gae/domain"
	"time"
)

type CommentUseCase struct {
	commentRepository   CommentRepository
	commenterRepository CommenterRepository
	pageRepository      PageRepository
}

func NewCommentUseCase(commentRepository CommentRepository, pageRepository PageRepository, commenterRepository CommenterRepository) *CommentUseCase {
	return &CommentUseCase{
		commentRepository:   commentRepository,
		commenterRepository: commenterRepository,
		pageRepository:      pageRepository,
	}
}

func (u *CommentUseCase) PostComment(url string, name string, commentText string) {
	// get or create page
	page := u.pageRepository.FindByUrl(url)
	if page == nil {
		page = domain.NewPage(u.pageRepository.NextPageId(), url)
	}

	commenter := domain.NewCommenter(u.commenterRepository.NextCommenterId(), name)

	comment := commenter.NewComment(commentText, page, time.Now())
	u.commentRepository.Add(comment)
}

func (u *CommentUseCase) GetComments(url string) []*domain.Comment {
	page := u.pageRepository.FindByUrl(url)
	if page == nil {
		return []*domain.Comment{}
	}
	comments := u.commentRepository.FindByPageId(page.PageId())
	return comments
}

type Repository interface {
}

type CommentRepository interface {
	Repository
	NextCommentId() domain.CommentId
	Add(comment *domain.Comment)
	Delete(comment *domain.Comment)
	FindByPageId(page domain.PageId) []*domain.Comment
}

type PageRepository interface {
	Repository
	NextPageId() domain.PageId
	Add(post *domain.Comment)
	Delete(post *domain.Comment)
	FindByUrl(url string) *domain.Page
}

type CommenterRepository interface {
	Repository
	NextCommenterId() domain.CommenterId
	Add(post *domain.Commenter)
	Delete(post *domain.Commenter)
	FindById(page *domain.Page) *domain.Commenter
}
