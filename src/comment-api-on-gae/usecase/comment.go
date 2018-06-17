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

func (u *CommentUseCase) PostNewComment(url string, name string, commentText string) {
	page := u.pageRepository.FindByUrl(url)
	if page == nil {
		page = domain.NewPage(u.pageRepository.NextPageId(), url)
	}

	commenter := domain.NewCommenter(u.commenterRepository.NextCommenterId(), name)

	comment := commenter.CreateNewComment(commentText, page, time.Now())
	u.commentRepository.Add(comment)
}

type Repository interface {
}


type CommentRepository interface {
	Repository
	NextCommentId() *domain.CommentId
	Add(comment *domain.Comment)
	Delete(comment *domain.Comment)
	FindByPage(page *domain.Page)
}

type PageRepository interface {
	Repository
	NextPageId() *domain.PageId
	Add(post *domain.Comment)
	Delete(post *domain.Comment)
	FindByUrl(url string) *domain.Page
}

type CommenterRepository interface {
	Repository
	NextCommenterId() *domain.CommenterId
	Add(post *domain.Commenter)
	Delete(post *domain.Commenter)
	FindById(page *domain.Page) *domain.Commenter
}
