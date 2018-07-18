package usecase

import (
	"commenting/domain"
	"time"
)

type CommentUseCase struct {
	commentRepository   CommentRepository
	commenterRepository CommenterRepository
	pageRepository      PageRepository
}

// あんまりいらない気もするけどデメリットもない気がするのでとりあえず作っている
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
	if !domain.IsValidPageId(strPageId) {
		return &Result{
			code:    E_INVALID,
			message: "PageId is invalid",
		}
	}
	pageId := domain.NewPageId(strPageId)

	// TODO: 以下はsnippet化したくなりそう
	// usecaseに関してはDRYじゃなくても弊害少ない？
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
	if !domain.IsValidPageId(strPageId) {
		return nil, nil, &Result{
			// messageがDRYじゃないけどそんなに弊害無いと判断、messageの時点でそもそも統一性持たせなくていい前提
			// 統一性もたせる必要があるならcodeをそのレベルまで細分化すべき
			message: "PageId is invalid",
			code:    E_INVALID,
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
