package main

import (
	"comment-api-on-gae/middleware"
	"net/http"
)

func init() {
	http.Handle("/", middleware.NewHandler())
}
