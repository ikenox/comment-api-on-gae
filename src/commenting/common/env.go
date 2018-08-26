package env

import (
	"os"
	"strings"
	"time"
)

var CurrentTime = time.Now

var isRunOnGAE = func() bool {
	// https://cloud.google.com/appengine/docs/standard/python/tools/using-local-server#detecting_application_runtime_environment
	return strings.HasPrefix(os.Getenv("SERVER_SOFTWARE"), "Google App Engine/")
}()

func IsProduction() bool {
	return isRunOnGAE
}

func IsDevelopment() bool {
	return !isRunOnGAE
}

var Namespace string = func() string {
	// 複数サーバーで動いているときなどnamespaceで区別
	if ns := os.Getenv("NAMESPACE"); ns != "" {
		return ns
	}
	return "app"
}()
