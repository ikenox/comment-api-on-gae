package usecase

type Code string

const (
	EINVALID  Code = "invalid"   // validation failed
	ENOTFOUND Code = "not found" // validation failed
	EINTERNAL Code = "internal"  // unknown error
)

// レイヤごとにエラー型を別に定義している
// - 想定外のエラーがレイヤ超えて飛んでくることがなくなる
// - 役割が異なる。ひとつのドメインエラーに対してユーザーには複数の伝え方（アプリケーションエラー）があるかも
// - 何が起きたかはユーザーには詳細を伝えたくない場合もあるかもしれない
//   - ドメインエラー: 純粋に何が起きたかを厳密に表現
//   - アプリケーションエラー: ユーザーへの伝え方を表現
type Error struct {
	message string
	code    Code
}

func (e *Error) Code() Code {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}
