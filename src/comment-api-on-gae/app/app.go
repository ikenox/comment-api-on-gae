package app

import (
	"net/http"
)

func init() {
	// serve
	http.Handle("/", NewServer())
}
