package usecase

import (
	"comment-api-on-gae/commenting/domain"
	"comment-api-on-gae/common/usecase"
	"fmt"
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

	pageId := domain.NewPageID(strPageId)
	page := u.pageRepository.Get(pageId)
	if page == nil {
		// create new page if not exist
		page = domain.NewPage(pageId)
	}
	u.pageRepository.Add(page)

	userID := u.commenterRepository.CurrentUser(idToken)
	// name
	if name == "" {
		return nil, usecase.NewResult(usecase.INVALID, "name must not be empty.")
	}

	if util.LengthOf(name) > 20 {
		return nil, usecase.NewResult(usecase.INVALID, "name should be less than 20 characters.")
	}
	commenter := domain.NewCommenter(u.commenterRepository.NextCommenterID(), name, userID)
	u.commenterRepository.Put(commenter)

	comment := commenter.MakeComment(u.commentRepository.NextCommentID(), text, page, env.CurrentTime())
	u.commentRepository.Put(comment)

	// TODO イベントpublish部分の実装やインターフェースが雑
	// TODO 実践ドメイン駆動設計ではドメイン層からpublishしているがどうすべきか
	u.publisher.Publish("CommentPosted", fmt.Sprintf("name:%s;comment:%s;", name, text))
	u.log.Infof("label:CommentPosted,name:%s,comment:%s", name, text)

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

func (u *CommentUseCase) DeleteComment(idToken string, commentIDStr string) *usecase.Result {
	userID := u.commenterRepository.CurrentUser(idToken)
	if userID == "" {
		return usecase.NewResult(usecase.INVALID, "login required.")
	}

	commentIDInt, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		panic(err)
	}
	commentID := domain.CommentID(commentIDInt)
	comment := u.commentRepository.Get(commentID)
	//if userID != comment.commenter.userID {
	if comment != nil {
		return usecase.NewResult(usecase.INVALID, "not allowed.")
	}

	u.commentRepository.Delete(commentID)

	return usecase.NewResult(usecase.OK, "")
}
