package usecase

import (
	"commenting/domain"
	"fmt"
	"time"
	"unicode/utf8"
)

type CommentUseCase struct {
	commentRepository   CommentRepository
	commenterRepository CommenterRepository
	pageRepository      PageRepository
}

// あんまりいらない気もする
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

// デメリットあんまり無さそうなのでコマンドとクエリの責務分離してない
func (u *CommentUseCase) PostComment(strPageId string, name string, text string) *Result {
	// ドメイン層でPageIdのバリデーションエラーハンドリングしようとするといたるところにエラーハンドリングが散らばるのでやめた方良さそう
	// この例だと、NewPageIdがエラー返しちゃうとstring => PageIdの変換をするいたるところにエラーハンドリングのボイラープレートロジックが書かれる
	// ドメインのどこかで発生してたらい回しにされまくって返ってきたエラーをcaseに分けてハンドリングするの辛い
	// failure is your domainでは再帰的にエラーメッセージ探索してるけど、ドメイン層で発生したエラーのエラーメッセージってほんとにユーザーに見せていいの？
	// なるべく外側のレイヤ(アプリケーション層)でバリデーションする前提でドメインロジック書いたほうがドメインがスッキリしそう
	// 内側のレイヤなほどerror投げたときにそれをキャッチする処理かかなくてはいけなくなる箇所が増える
	// そもそもエラーはドメインの概念じゃない？
	// 実行時エラー返すんじゃなくて、「何が正常か」を明示的に表現している(メソッドが生えている)方がドメインモデルのあり方としては正しい気がする
	// ドメイン層は純粋なものしか扱わないようにする、不純物混ざりそうになった瞬間に即panicする方針
	// アプリケーション層以下では不純物混ざらないという前提で書けるので全体的に記述量減るしシンプルになる気がした
	// ドメインにIsValidXXといったメソッド増えまくりそうなのはちょっとあれかも。static method欲しくなる。。
	if err := domain.PageIdSpec.CheckValidityOf(strPageId); err != nil {
		return &Result{
			ErrInvalid,
			fmt.Sprintf("PageId is invalid: %s", err.Error()),
		}
	}

	if nameLen := utf8.RuneCountInString(name); nameLen > 20 {
		return &Result{ErrInvalid, "Name is too long."}
	}

	// TODO: ここらへんドメインで定義されるべきか
	if text == "" {
		return &Result{ErrInvalid, "Comment is too long."}
	}
	if commentLen := utf8.RuneCountInString(text); commentLen > 1000 {
		return &Result{ErrInvalid, "Comment is too long."}
	}

	pageId := domain.NewPageId(strPageId)
	page := u.pageRepository.Get(pageId)
	if page == nil {
		page = domain.NewPage(pageId)
	}
	u.pageRepository.Add(page)

	commenter := domain.NewCommenter(u.commenterRepository.NextCommenterId(), name)
	comment := commenter.NewComment(u.commentRepository.NextCommentId(), text, page, time.Now())

	u.commenterRepository.Add(commenter)
	u.commentRepository.Add(comment)

	return &Result{code: OK}
}

func (u *CommentUseCase) GetComments(strPageId string) ([]*domain.Comment, []*domain.Commenter, *Result) {
	if err := domain.PageIdSpec.CheckValidityOf(strPageId); err != nil {
		// messageがDRYじゃないけどそんなに弊害無いと判断、messageの時点でそもそも統一性持たせなくていい前提
		// 統一性もたせる必要があるならcodeをそのレベルまで細分化すべき
		return nil, nil, &Result{
			ErrInvalid,
			fmt.Sprintf("PageId is invalid: %s", err.Error()),
		}
	}
	pageId := domain.NewPageId(strPageId)

	page := u.pageRepository.Get(pageId)
	if page == nil {
		return []*domain.Comment{}, []*domain.Commenter{}, &Result{code: OK}
	}

	comments := u.commentRepository.FindByPageId(page.PageId())

	commentIds := make([]domain.CommenterId, len(comments))
	for i, c := range comments {
		commentIds[i] = c.CommenterId()
	}
	commenters := u.commenterRepository.FindByComments(commentIds)

	// TODO: もうちょっとかっこよく返したい
	// comment, commenterはどちらも存在するor存在しないというのをコードで表明できていない
	return comments, commenters, &Result{code: OK}
}
