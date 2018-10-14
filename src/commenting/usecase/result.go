package usecase

type Code string

// UseCaseの処理結果を表すコード
// クライアントに伝えたい単位で分割
// 結果を抽象化しておくことでcontrollerではhttp status code以外への変換にも対応できる
// この例ではhttp status codeと同じくらいの粒度だが、アプリの要件によってはもっと細分化されるかも
// なんのフィールドがバリデーションエラーになったかとかまで構造化してクライアントに知らせる必要あるならそのレベルまで細分化してエラーコード定義する必要ありそう
// Resultに結果のデータも持たせようと思ったがgolangだとジェネリクスがなくて型が失われるのでやめた
// ジェネリクスある言語なら持たせていい気がする
// 成功も失敗も本質的には変わらないのでは？
// TODO: 処理は失敗したけどなんらかのデータを返したい場合ってある？返さないほうがいい？
// TODO: OK, NGをrootとする階層構造にすればon demandで細分化できて良さそう？
// Result
//   OK
//     CREATED
//
//   Error
//     INVALID
//       PAGEID
const (
	OK         Code = "ok"         // success
	CREATED    Code = "created"    // create success
	INVALID    Code = "invalid"    // validation failed
	NOTFOUND   Code = "not found"  // resource not found
	UNEXPECTED Code = "unexpected" // unknown error
)

type Result struct {
	code    Code
	message string
}

func NewResult(code Code, message string) *Result {
	return &Result{
		code:    code,
		message: message,
	}
}

func (e *Result) Code() Code {
	return e.code
}

func (e *Result) Message() string {
	return e.message
}
