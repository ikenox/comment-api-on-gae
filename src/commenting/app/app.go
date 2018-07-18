package main

import (
	"commenting/middleware"
	"net/http"
)

func init() {
	http.Handle("/", middleware.NewHandler())
}
