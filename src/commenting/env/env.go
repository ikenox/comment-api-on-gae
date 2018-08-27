package env

import (
	"firebase.google.com/go"
	"time"
)

// application globals
// アプリ起動以降不変とみなせる値や関数、環境変数など？
// レイヤ関係なく使われるような普遍的な値や概念はRepositoryにするよりグローバル変数にしてしまったほうが良さそう
var Namespace string = "app"
var FirebaseApp *firebase.App = nil
var CurrentTime func() time.Time = time.Now
var IsProduction bool = false
