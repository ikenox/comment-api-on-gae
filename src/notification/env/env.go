package env

import (
	"os"
)

// application globals
// アプリにとって普遍とみなせる値や関数、環境変数など？
// どこからでも使われるような値や概念はRepositoryにするよりグローバル変数にしてしまったほうが楽そう
var Namespace string
var IsProduction bool
var ProjectID string
var AdminEmail string
var SenderEmail string

func init() {
	ProjectID = os.Getenv("APP_ID")

	IsProduction = os.Getenv("IS_PRODUCTION") == "True"

	AdminEmail = os.Getenv("ADMIN_EMAIL")
	SenderEmail = os.Getenv("SENDER_EMAIL")

	// namespace
	if ns := os.Getenv("NAMESPACE"); ns != "" {
		Namespace = ns
	} else {
		Namespace = ""
	}
}
