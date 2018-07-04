package usecase

type Code string

// UseCaseの処理結果全部ここ
// 成功も失敗も本質的には変わらないのでは？
// TODO 考察: 処理は失敗したけどなんらかのデータを返したい場合ってある？返さないほうがいい？
// 実質http status codeと1対多の関係？
// 結果を抽象化しておくことでcontrollerではhttp status code以外への変換にも対応できる
const (
	OK          Code = "ok"         // success
	CREATED     Code = "created"    //
	EINVALID    Code = "invalid"    // validation failed
	ENOTFOUND   Code = "not found"  // validation failed
	EUNEXPECTED Code = "unexpected" // unknown error
)

type Result struct {
	data    interface{}
	message string
	code    Code
}

func (e *Result) Code() Code {
	return e.code
}

func (e *Result) Message() string {
	return e.message
}
