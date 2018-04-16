package main

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", commentHandler.List)
}
