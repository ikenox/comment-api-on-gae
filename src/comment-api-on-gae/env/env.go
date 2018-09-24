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
// レイヤ関係なく使われるような普遍的な値や概念はRepositoryにするよりグローバル変数にしてしまったほうが良さそう
var Namespace string
var CurrentTime func() time.Time
var IsProduction bool
var FirebaseApp *firebase.App

func init() {
	ctx := context.Background()

	// time
	CurrentTime = time.Now

	// namespace
	if ns := os.Getenv("NAMESPACE"); ns != "" {
		Namespace = ns
	} else {
		Namespace = "default"
	}

	// firebase
	path, err := filepath.Abs(os.Getenv("FIREBASE_SERVICE_ACCOUNT_PATH"))
	if err != nil {
		panic(err.Error())
	}
	opt := option.WithCredentialsFile(path)
	if app, err := firebase.NewApp(ctx, nil, opt); err == nil {
		FirebaseApp = app
	} else {
		panic(err.Error())
	}

	// dev or prod
	// https://cloud.google.com/appengine/docs/standard/python/tools/using-local-server#detecting_application_runtime_environment
	IsProduction = os.Getenv("IS_PRODUCTION") == "true"
}
