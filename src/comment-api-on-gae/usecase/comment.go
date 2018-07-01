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

func (u *CommentUseCase) PostComment(strPageId string, name string, text string) *Error {
	pageId, err := domain.NewPageId(strPageId)
	// TODO: 想定外のエラーへの対応
	if err != nil {
		return &Error{
			domainError: err,
			code:        EINVALID,
		}
	}

	// get or create page
	// TODO: 以下はsnippet化したくなりそう
	// usecaseに関してはDRYじゃなくても弊害少ないか？
	// コード的には重複していても概念的には別のケースを表すコードになる？
	page := u.pageRepository.Get(pageId)
	if page == nil {
		page = domain.NewPage(u.pageRepository.NextPageId())
	}
	u.pageRepository.Add(page)

	commenter := domain.NewCommenter(u.commenterRepository.NextCommenterId(), name)
	u.commenterRepository.Add(commenter)

	comment := commenter.NewComment(u.commentRepository.NextCommentId(), text, page, time.Now())
	u.commentRepository.Add(comment)
	return nil
}

func (u *CommentUseCase) GetComments(id string) ([]*domain.Comment, *Error) {
	pageId, err := domain.NewPageId(id)
	// TODO: 想定外のエラーへの対応
	if err != nil {
		return nil, &Error{
			domainError: err,
			code:        EINVALID,
		}
	}

	page := u.pageRepository.Get(pageId)
	if page == nil {
		return nil, &Error{
			code: ENOTFOUND,
		}
	}
	comments := u.commentRepository.FindByPageId(page.PageId())
	return comments, nil
}

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
	Add(commenter *domain.Commenter)
	Delete(commenterId domain.CommenterId)
	Get(commenterId domain.CommenterId) *domain.Commenter
}
