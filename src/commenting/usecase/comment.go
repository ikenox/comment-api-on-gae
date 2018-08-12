package usecase

import (
	"commenting/domain"
	"commenting/usecase/validator"
	"fmt"
	"time"
)

type CommentUseCase struct {
	commentRepository   CommentRepository
	commenterRepository CommenterRepository
	pageRepository      PageRepository
}

func NewCommentUseCase(
	commentRepo CommentRepository,
	commenterRepo CommenterRepository,
	pageRepo PageRepository,
) *CommentUseCase {
	return &CommentUseCase{
		commentRepository:   commentRepo,
		commenterRepository: commenterRepo,
		pageRepository:      pageRepo,
	}
}

type CommentWithCommenter struct {
	Comment   *domain.Comment
	Commenter *domain.Commenter
}

func (u *CommentUseCase) PostComment(strPageId string, name string, text string) (*CommentWithCommenter, *Result) {
	// ドメイン層でPageIdのバリデーションエラーハンドリングしようとするといたるところにエラーハンドリングが散らばるので極力避けたほういい？
	// この例だと、NewPageIdがエラー返しちゃうとstring => PageIdの変換をするいたるところにエラーハンドリングのボイラープレートロジックが書かれる
	// 内側のレイヤなほどerror投げたときにそれをキャッチする処理かかなくてはいけなくなる箇所が増える
	// ドメインのどこかで発生してたらい回しにされまくって返ってきたエラーをcaseに分けてハンドリングするの辛い
	// それは本当にドメインロジックなのかどうかを考えるのが良さそう（e.g. コメントの文字数制限したいのはアプリケーションの都合）
	// なるべく外側のレイヤ(アプリケーション層)でバリデーションする前提でドメインロジック書いたほうがドメインがスッキリしそう
	// 全体的に記述量減るしシンプルになる気がした
	// そもそもエラーはドメインの概念じゃない？
	// 実行時エラー返すんじゃなくて、「何が正常か」を明示的に表現している(メソッドが生えている)方がドメインモデルのあり方としては正しい気がする
	if err := validator.ValidatePageID(strPageId); err != nil {
		return nil, &Result{
			ErrInvalid,
			fmt.Sprintf(err.Error()),
		}
	}
	if err := validator.ValidateComment(text); err != nil {
		return nil, &Result{
			ErrInvalid,
			fmt.Sprintf(err.Error()),
		}
	}
	if err := validator.ValidateCommenterName(name); err != nil {
		return nil, &Result{
			ErrInvalid,
			fmt.Sprintf(err.Error()),
		}
	}

	pageId := domain.NewPageID(strPageId)
	page := u.pageRepository.Get(pageId)
	if page == nil {
		page = domain.NewPage(pageId)
	}
	u.pageRepository.Add(page)

	commenter := domain.NewCommenter(u.commenterRepository.NextCommenterID(), name)
	comment := commenter.NewComment(u.commentRepository.NextCommentID(), text, page, time.Now())

	u.commenterRepository.Add(commenter)
	u.commentRepository.Add(comment)

	// デメリットあんまり無さそうなのでコマンドとクエリの責務分離してない
	return &CommentWithCommenter{comment, commenter}, &Result{code: OK}
}

func (u *CommentUseCase) ReplyComment(commentID int64, name string, text string) (*CommentWithCommenter, *Result) {
	if err := validator.ValidateComment(text); err != nil {
		return nil, &Result{
			ErrInvalid,
			fmt.Sprintf(err.Error()),
		}
	}
	if err := validator.ValidateCommenterName(name); err != nil {
		return nil, &Result{
			ErrInvalid,
			fmt.Sprintf(err.Error()),
		}
	}

	commenter := domain.NewCommenter(u.commenterRepository.NextCommenterID(), name)
	comment := commenter.NewComment(u.commentRepository.NextCommentID(), text, page, time.Now())

	u.commenterRepository.Add(commenter)
	u.commentRepository.Add(comment)

	// デメリットあんまり無さそうなのでコマンドとクエリの責務分離してない
	return &CommentWithCommenter{comment, commenter}, &Result{code: OK}
}

func (u *CommentUseCase) GetComments(strPageID string) ([]*CommentWithCommenter, *Result) {
	pageId := domain.NewPageID(strPageID)

	page := u.pageRepository.Get(pageId)
	if page == nil {
		return []*CommentWithCommenter{}, &Result{code: OK}
	}

	comments := u.commentRepository.FindByPageID(page.PageId())

	commentIds := make([]domain.CommenterID, len(comments))
	for i, c := range comments {
		commentIds[i] = c.CommenterId()
	}
	commenters := u.commenterRepository.FindByComments(commentIds)

	data := make([]*CommentWithCommenter, len(comments))
	if len(comments) > 0 {
		for i := 0; i < len(comments); i++ {
			if commenters[i] != nil && comments[i] != nil {
				data[i] = &CommentWithCommenter{
					Comment:   comments[i],
					Commenter: commenters[i],
				}
			}
		}
	}

	return data, &Result{code: OK}
}
