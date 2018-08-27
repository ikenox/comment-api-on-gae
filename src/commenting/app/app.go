package app

import (
	"commenting/env"
	"commenting/middleware"
	"context"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"google.golang.org/appengine/log"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	ctx := context.Background()

	// time
	env.CurrentTime = func() time.Time { return time.Now() }

	// namespace
	if ns := os.Getenv("NAMESPACE"); ns != "" {
		env.Namespace = ns
	}

	// firebase
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	if app, err := firebase.NewApp(ctx, nil, opt); err == nil {
		env.FirebaseApp = app
	} else {
		log.Errorf(ctx, err.Error())
	}

	// dev or prod
	// https://cloud.google.com/appengine/docs/standard/python/tools/using-local-server#detecting_application_runtime_environment
	env.IsProduction = strings.HasPrefix(os.Getenv("SERVER_SOFTWARE"), "Google App Engine/")

	// serve
	http.Handle("/", middleware.NewServer())
}
