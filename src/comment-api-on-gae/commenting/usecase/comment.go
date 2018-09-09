package usecase

import (
	"comment-api-on-gae/commenting/domain"
	"comment-api-on-gae/common/usecase"
	"comment-api-on-gae/env"
	"comment-api-on-gae/util"
	"regexp"
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

func (u *CommentUseCase) PostComment(idToken string, name string, strPageId string, text string) (*CommentWithCommenter, *usecase.Result) {
	// ドメイン層でPageIdのバリデーションエラーハンドリングしようとするといたるところにエラーハンドリングが散らばるので極力避けたほういい？
	// この例だと、NewPageIdがエラー返しちゃうとstring => PageIdの変換をするいたるところにエラーハンドリングのボイラープレートロジックが書かれる
	// 内側のレイヤなほどerror投げたときにそれをキャッチする処理かかなくてはいけなくなる箇所が増える、永続層から取り出すときとかにもバリデーションが走ることになる
	// それは本当にドメインロジックなのかどうかを考えるのが良さそう（e.g. コメントの文字数制限したいのはアプリケーションの都合）
	// なるべく外側のレイヤ(アプリケーション層)でバリデーションする前提でドメインロジック書いたほうがドメインがスッキリしそう
	// 全体的に記述量減るしシンプルになる気がした
	// そもそも「エラー」はドメインの概念じゃない？

	// pageId
	if strPageId == "" {
		return nil, usecase.NewResult(usecase.INVALID, "PageID must not be empty")
	}
	if strPageId = pageIdRegexp.Copy().FindString(strPageId); strPageId == "" {
		return nil, usecase.NewResult(usecase.INVALID, "invalid character")
	}
	if len(strPageId) > 64 {
		return nil, usecase.NewResult(usecase.INVALID, "page ID is too long")
	}

	// text
	if text == "" {
		return nil, usecase.NewResult(usecase.INVALID, "comment must not be empty")
	}
	if util.LengthOf(text) > 1000 {
		return nil, usecase.NewResult(usecase.INVALID, "comment is too long")
	}

	pageId := domain.NewPageID(strPageId)
	page := u.pageRepository.Get(pageId)
	if page == nil {
		// create new page if not exist
		page = domain.NewPage(pageId)
	}
	u.pageRepository.Add(page)

	commenter := u.commenterRepository.CurrentCommenter(idToken)
	if commenter == nil {
		// name
		if util.LengthOf(name) > 20 {
			return nil, usecase.NewResult(usecase.INVALID, "commenter name is too long")
		}
		commenter = domain.NewCommenter(u.commenterRepository.NextCommenterID(), name, "")
		u.commenterRepository.Put(commenter)
	}
	comment := commenter.MakeComment(u.commentRepository.NextCommentID(), text, page, env.CurrentTime())

	u.commentRepository.Put(comment)

	return &CommentWithCommenter{comment, commenter}, usecase.NewResult(usecase.OK, "")
}

func (u *CommentUseCase) GetComments(strPageID string) ([]*CommentWithCommenter, *usecase.Result) {
	pageId := domain.NewPageID(strPageID)

	page := u.pageRepository.Get(pageId)
	if page == nil {
		return []*CommentWithCommenter{}, usecase.NewResult(usecase.OK, "")
	}

	comments := u.commentRepository.FindByPageID(page.PageId())

	userIDs := make([]domain.CommenterID, len(comments))
	for i, c := range comments {
		userIDs[i] = c.CommenterID()
	}
	commenters := u.commenterRepository.FindByCommenterID(userIDs)

	data := make([]*CommentWithCommenter, len(comments))
	if len(comments) > 0 {
		for i := 0; i < len(comments); i++ {
			data[i] = &CommentWithCommenter{
				Comment:   comments[i],
				Commenter: commenters[i],
			}
		}
	}

	return data, usecase.NewResult(usecase.OK, "")
}
