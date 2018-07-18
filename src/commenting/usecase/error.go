package usecase

type Code string

// UseCaseの処理結果を表すコード
// ユーザーに伝えたい単位で分割
// 結果を抽象化しておくことでcontrollerではhttp status code以外への変換にも対応できる
// この例ではhttp status codeと同じくらいの粒度だが、アプリの要件によってはもっと細分化されるかも
// なんのフィールドがバリデーションエラーになったかとかまでクライアントに知らせる必要あるならそのレベルでエラーコード定義
// 成功も失敗も本質的には変わらないのでは？
// Resultに結果のデータも持たせようと思ったがgolangだとジェネリクスがなくて型が失われるのでやめた
// ジェネリクスある言語なら持たせていい気がする
// TODO: 処理は失敗したけどなんらかのデータを返したい場合ってある？返さないほうがいい？
// TODO: 同じエラーがいろんなとこでinstanciate
// TODO: 階層構造にすればon demandで細分化できて良さそう？？
// Result
//   OK
//     CREATED
//
//   Error
//     INVALID
//       PAGEID
const (
	OK          Code = "ok"         // success
	E_INVALID   Code = "invalid"    // validation failed
	E_NOTFOUND  Code = "not found"  // validation failed
	E_NEXPECTED Code = "unexpected" // unknown error
)

// データの入れ物でしかないのでコンストラクタ不要と判断
type Result struct {
	code    Code
	message string
}

func (e *Result) Code() Code {
	return e.code
}

func (e *Result) Message() string {
	return e.message
}
