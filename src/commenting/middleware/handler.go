package middleware

import (
	"commenting/common"
	"commenting/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func NewServer() http.Handler {
	s := standard.New("")
	s.SetHandler(NewEcho())
	return s
}

func NewEcho() engine.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	//e.Use(middleware.Recover())
	e.Use(useAppEngine)

	pc := controller.NewCommentController()
	e.GET("/comment", pc.List)
	e.POST("/comment", pc.PostComment)
	return e
}

func useAppEngine(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if r, ok := c.Request().(*standard.Request); ok {
			namespace := env.Namespace
			ctx := appengine.WithContext(c.StdContext(), r.Request)
			ctx, err := appengine.Namespace(ctx, namespace)
			if err != nil {
				log.Errorf(ctx, "unresolve to set namespace (err %v)", err)
			}
			c.SetStdContext(ctx)
		}
		return next(c)
	}
}
