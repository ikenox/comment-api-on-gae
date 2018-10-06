package env

import (
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"os"
	"path/filepath"
	"time"
)

// application globals
// アプリにとって普遍とみなせる値や関数、環境変数など？
// レイヤ関係なく使われるような普遍的な値や概念はRepositoryにするよりグローバル変数にしてしまったほうが楽そう
var Namespace string
var CurrentTime func() time.Time
var IsProduction bool
var FirebaseApp *firebase.App
var GCPCredentialOption option.ClientOption
var ProjectID string

func init() {
	ctx := context.Background()

	// FIXME
	ProjectID = "comment-api-dev"

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

	// FIXME クライアント郡がenvにいるのはおかしいので、optionだけenvに残して移動

	// firebase
	if FirebaseApp, err = firebase.NewApp(ctx, nil, GCPCredentialOption); err != nil {
		panic(err.Error())
	}
}
