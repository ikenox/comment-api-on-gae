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
	// ドメイン層でPageIdのバリデーションエラーハンドリングしようとするといたるところにエラーハンドリングが散らばるのでやめた方良さそう
	// この例だと、NewPageIdがエラー返しちゃうとstring => PageIdの変換をするいたるところにエラーハンドリングのボイラープレートロジックが書かれる
	// ドメインのどこかで発生してたらい回しにされまくって返ってきたエラーをcaseに分けてハンドリングするの辛い
	// なるべく外側のレイヤ(アプリケーション層)でバリデーションする前提でドメインロジック書いたほうがドメインがスッキリしそう
	// 内側のレイヤなほどerror投げたときにそれをキャッチする処理かかなくてはいけなくなる箇所が増える
	// そもそもエラーはドメインの概念じゃない？
	// 実行時エラー返すんじゃなくて、「何が正常か」を明示的に表現している(メソッドが生えている)方がドメインモデルのあり方としては正しい気がする
	// ドメイン層は純粋なものしか扱わないようにする、不純物混ざりそうになった瞬間に即panicする方針
	// アプリケーション層以下では不純物混ざらないという前提で書けるので全体的に記述量減るしシンプルになる気がした
	// ドメインにIsValidXXといったメソッド増えまくりそうなのはちょっとあれかも。static method欲しくなる。。
	if !domain.IsValidPageId(strPageId) {
		return &Error{
			message: "PageId is invalid",
			code:    EINVALID,
		}
	}
	pageId := domain.NewPageId(strPageId)

	// TODO: 以下はsnippet化したくなりそう
	// usecaseに関してはDRYじゃなくても弊害少ない？
	// Get or Create Page
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

func (u *CommentUseCase) GetComments(strPageId string) ([]*domain.Comment, *Error) {
	if !domain.IsValidPageId(strPageId) {
		return nil, &Error{
			message: "PageId is invalid",
			code:    EINVALID,
		}
	}
	pageId := domain.NewPageId(strPageId)

	page := u.pageRepository.Get(pageId)
	if page == nil {
		return nil, &Error{
			message: "Page is not found",
			code:    ENOTFOUND,
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
