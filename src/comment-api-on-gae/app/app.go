package main

import (
	"comment-api-on-gae/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(UseAppEngine)

	pageController := controller.NewCommentController()
	e.GET("/comment", pageController.List)
	e.POST("/comment", pageController.PostComment)

	s := standard.New("")
	s.SetHandler(e)
	http.Handle("/", s)
}

func UseAppEngine(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if r, ok := c.Request().(*standard.Request); ok {
			namespace := "development"
			ctx := appengine.WithContext(c.StdContext(), r.Request)
			ctx, err := appengine.Namespace(ctx, namespace)
			if err != nil {
				log.Errorf(ctx, "unresolve to set namespace (err %v)", err)
			}
			log.Infof(ctx, "namespace:%s", namespace)
			c.SetStdContext(ctx)
		}
		return next(c)
	}
}

type Context struct {
	ctx echo.Context
}
