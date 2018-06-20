package main

import (
	"comment-api-on-gae/controller"
	"net/http"

	"github.com/labstack/echo"
)

func init() {
	pageController := controller.NewPageController()
	http.HandleFunc("/comment/list", pageController.List)
	http.HandleFunc("/comment/add", pageController.Add)
	http.ListenAndServe(":8080", nil)
}
