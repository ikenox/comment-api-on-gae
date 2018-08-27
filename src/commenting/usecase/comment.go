package usecase

import (
	"commenting/domain"
	"commenting/env"
	"regexp"
	"util"
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

var pageIdRegexp = regexp.MustCompile("^[0-9a-zA-Z_\\-]+$")

func (u *CommentUseCase) PostComment(strPageId string, name string, text string) (*CommentWithCommenter, *Result) {
	// ドメイン層でPageIdのバリデーションエラーハンドリングしようとするといたるところにエラーハンドリングが散らばるので極力避けたほういい？
	// この例だと、NewPageIdがエラー返しちゃうとstring => PageIdの変換をするいたるところにエラーハンドリングのボイラープレートロジックが書かれる
	// 内側のレイヤなほどerror投げたときにそれをキャッチする処理かかなくてはいけなくなる箇所が増える、永続層から取り出すときとかにもバリデーションが走ることになる
	// それは本当にドメインロジックなのかどうかを考えるのが良さそう（e.g. コメントの文字数制限したいのはアプリケーションの都合）
	// なるべく外側のレイヤ(アプリケーション層)でバリデーションする前提でドメインロジック書いたほうがドメインがスッキリしそう
	// 全体的に記述量減るしシンプルになる気がした
	// そもそも「エラー」はドメインの概念じゃない？

	// pageId
	if strPageId == "" {
		return nil, &Result{
			code:    INVALID,
			message: "PageID must not be empty",
		}
	}
	if strPageId = pageIdRegexp.Copy().FindString(strPageId); strPageId == "" {
		return nil, &Result{
			code:    INVALID,
			message: "invalid character",
		}
	}
	if len(strPageId) > 64 {
		return nil, &Result{
			code:    INVALID,
			message: "page ID is too long",
		}
	}

	// text
	if text == "" {
		return nil, &Result{
			code:    INVALID,
			message: "comment must not be empty",
		}
	}
	if util.LengthOf(text) > 1000 {
		return nil, &Result{
			code:    INVALID,
			message: "comment is too long",
		}
	}

	// name
	if util.LengthOf(name) > 20 {
		return nil, &Result{
			code:    INVALID,
			message: "commenter name is too long",
		}
	}

	pageId := domain.NewPageID(strPageId)
	page := u.pageRepository.Get(pageId)
	if page == nil {
		// create new page if not exist
		page = domain.NewPage(pageId)
	}
	u.pageRepository.Add(page)

	commenter := domain.NewCommenter(u.commenterRepository.NextCommenterID(), name)
	comment := commenter.NewComment(u.commentRepository.NextCommentID(), text, page, env.CurrentTime())

	u.commenterRepository.Add(commenter)
	u.commentRepository.Add(comment)

	// コマンドとクエリの責務分離してないけどデメリットよりメリットが大きいと判断(複数回API叩かなくて良い)
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
