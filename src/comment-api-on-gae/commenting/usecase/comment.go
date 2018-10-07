package usecase

import (
	"comment-api-on-gae/commenting/domain"
	"comment-api-on-gae/common/usecase"
	"strconv"

	"comment-api-on-gae/env"
	"comment-api-on-gae/util"
	"regexp"
)

type CommentUseCase struct {
	commentRepository   CommentRepository
	commenterRepository CommenterRepository
	pageRepository      PageRepository
	publisher           EventPublisher
	log                 LoggingRepository
}

func NewCommentUseCase(
	commentRepo CommentRepository,
	commenterRepo CommenterRepository,
	pageRepo PageRepository,
	publisher EventPublisher,
	log LoggingRepository,
) *CommentUseCase {
	return &CommentUseCase{
		commentRepository:   commentRepo,
		commenterRepository: commenterRepo,
		pageRepository:      pageRepo,
		publisher:           publisher,
		log:                 log,
	}
}

var pageIdRegexp = regexp.MustCompile("^[0-9a-zA-Z_\\-]+$")

func (u *CommentUseCase) PostComment(idToken string, name string, strPageId string, text string) (*domain.Comment, *usecase.Result) {
	// ドメイン層でPageIdのバリデーションエラーハンドリングしようとするといたるところにエラーハンドリングが散らばるので極力避けたほういい？
	// この例だと、NewPageIdがエラー返しちゃうとstring => PageIdの変換をするいたるところにエラーハンドリングのボイラープレートロジックが書かれる
	// 内側のレイヤなほどerror投げたときにそれをキャッチする処理かかなくてはいけなくなる箇所が増える、永続層から取り出すときとかにもバリデーションが走ることになる
	// それは本当にドメインロジックなのかどうかを考えるのが良さそう（e.g. コメントの文字数制限したいのはアプリケーションの都合）
	// なるべく外側のレイヤ(アプリケーション層)でバリデーションする前提でドメインロジック書いたほうがドメインがスッキリしそう
	// 全体的に記述量減るしシンプルになる気がした
	// そもそも「エラー」はドメインの概念じゃない？

	// pageId
	if strPageId == "" {
		return nil, usecase.NewResult(usecase.INVALID, "page id must not be empty.")
	}
	if strPageId = pageIdRegexp.Copy().FindString(strPageId); strPageId == "" {
		return nil, usecase.NewResult(usecase.INVALID, "page id contains invalid character.")
	}
	if len(strPageId) > 64 {
		return nil, usecase.NewResult(usecase.INVALID, "page id is too long.")
	}

	// text
	if text == "" {
		return nil, usecase.NewResult(usecase.INVALID, "comment must not be empty.")
	}
	if util.LengthOf(text) > 1000 {
		return nil, usecase.NewResult(usecase.INVALID, "comment should be less than 1000 characters.")
	}

	// name
	if name == "" {
		return nil, usecase.NewResult(usecase.INVALID, "name must not be empty.")
	}

	if util.LengthOf(name) > 20 {
		return nil, usecase.NewResult(usecase.INVALID, "name should be less than 20 characters.")
	}

	pageId := domain.NewPageID(strPageId)
	page := u.pageRepository.Get(pageId)
	if page == nil {
		// create new page if not exist
		page = domain.NewPage(pageId)
	}
	u.pageRepository.Add(page)

	commenter := u.commenterRepository.CurrentCommenter(idToken)
	comment := commenter.CreateComment(u.commentRepository.NextCommentID(), page.PageID(), text, name, env.CurrentTime())
	u.commentRepository.Put(comment)

	// TODO イベントpublish部分の実装やインターフェースが雑
	// TODO 実践ドメイン駆動設計ではドメイン層からpublishしているがどうすべきか
	// TODO きちんとイベントタイプごとにstructを定義してinterface層でmessageに変換
	u.publisher.Publish("CommentPosted", struct {
		// TODO jsonへのマッピングはinterface層の責務？
		CommentID int64  `json:"commentId"`
		PageID    string `json:"pageId"`
		Name      string `json:"name"`
		Text      string `json:"text"`
	}{
		CommentID: int64(comment.CommentID()),
		Name:      comment.Name(),
		Text:      comment.Text(),
		PageID:    string(comment.PageID()),
	})
	// TODO アプリケーションログのフォーマットのベタープラクティス
	u.log.Infof("label:CommentPosted,name:%s,comment:%s", name, text)

	return comment, usecase.NewResult(usecase.OK, "")
}

func (u *CommentUseCase) GetComments(strPageID string) ([]*domain.Comment, *usecase.Result) {
	pageId := domain.NewPageID(strPageID)

	page := u.pageRepository.Get(pageId)
	if page == nil {
		return []*domain.Comment{}, usecase.NewResult(usecase.OK, "")
	}

	comments := u.commentRepository.FindByPageID(page.PageID())
	return comments, usecase.NewResult(usecase.OK, "")
}

func (u *CommentUseCase) DeleteComment(idToken string, commentIDStr string) *usecase.Result {
	commenter := u.commenterRepository.CurrentCommenter(idToken)

	commentIDInt, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		panic(err)
	}
	commentID := domain.CommentID(commentIDInt)
	comment := u.commentRepository.Get(commentID)
	if comment == nil {
		return usecase.NewResult(usecase.NOTFOUND, "not found")
	}
	if commenter.UserID() != comment.Commenter().UserID() {
		return usecase.NewResult(usecase.INVALID, "not allowed.")
	}
	u.commentRepository.Delete(commentID)

	u.publisher.Publish("CommentDeleted", struct {
		CommentID int64  `json:"commentId"`
		PageID    string `json:"pageId"`
		Name      string `json:"name"`
		Text      string `json:"text"`
	}{
		CommentID: int64(comment.CommentID()),
		Name:      comment.Name(),
		Text:      comment.Text()[0:100],
		PageID:    string(comment.PageID()),
	})

	return usecase.NewResult(usecase.OK, "")
}
