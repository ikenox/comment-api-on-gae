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

func (u *CommentUseCase) PostComment(url string, name string, text string) {
	// get or create page
	// TODO: 以下はsnippet化したくなりそう
	// usecaseに関してはDRYじゃなくても弊害少ないか？
	// コード的には重複していても概念的には別のケースを表すコードになる？
	pageUrl := domain.NewPageUrl(url)
	page := u.pageRepository.FindByUrl(pageUrl)
	if page == nil {
		page = domain.NewPage(u.pageRepository.NextPageId(), pageUrl)
	}
	u.pageRepository.Add(page)

	commenter := domain.NewCommenter(u.commenterRepository.NextCommenterId(), name)
	u.commenterRepository.Add(commenter)

	comment := commenter.NewComment(u.commentRepository.NextCommentId(), text, page, time.Now())
	u.commentRepository.Add(comment)
}

func (u *CommentUseCase) GetComments(url string) []*domain.Comment {
	page := u.pageRepository.FindByUrl(domain.NewPageUrl(url))
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
	Delete(comment domain.CommentId)
	FindByPageId(page domain.PageId) []*domain.Comment
}

type PageRepository interface {
	Repository
	NextPageId() domain.PageId
	Add(page *domain.Page)
	Delete(page domain.PageId)
	FindByUrl(url *domain.PageUrl) *domain.Page
}

type CommenterRepository interface {
	Repository
	NextCommenterId() domain.CommenterId
	Add(commenter *domain.Commenter)
	Delete(commenterId domain.CommenterId)
	Get(commenterId domain.CommenterId) *domain.Commenter
}
