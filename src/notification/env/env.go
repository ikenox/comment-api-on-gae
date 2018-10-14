package env

import (
	"google.golang.org/api/option"
	"os"
	"path/filepath"
	"time"
)

// application globals
// アプリにとって普遍とみなせる値や関数、環境変数など？
// どこからでも使われるような値や概念はRepositoryにするよりグローバル変数にしてしまったほうが楽そう
var Namespace string
var CurrentTime func() time.Time
var IsProduction bool
var GCPCredentialOption option.ClientOption
var ProjectID string

func init() {
	ProjectID = os.Getenv("APP_ID")

	IsProduction = os.Getenv("IS_PRODUCTION") == "True"

	// time
	CurrentTime = time.Now

	// namespace
	if ns := os.Getenv("NAMESPACE"); ns != "" {
		Namespace = ns
	} else {
		Namespace = ""
	}

	path, err := filepath.Abs(os.Getenv("SERVICE_ACCOUNT_PATH"))
	if err != nil {
		panic(err.Error())
	}
	GCPCredentialOption = option.WithCredentialsFile(path)
}
