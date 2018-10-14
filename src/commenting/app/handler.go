package app

import (
	"commenting/env"
	"commenting/interface/controller"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/color"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"runtime"
)

func NewServer() http.Handler {
	s := standard.New("")
	s.SetHandler(NewEcho())
	return s
}

func NewEcho() engine.Handler {
	e := echo.New()

	e.SetHTTPErrorHandler(func(e error, context echo.Context) {
		stackSize := 4 << 10
		stack := make([]byte, stackSize)
		length := runtime.Stack(stack, true)
		log.Errorf(context.StdContext(), "[%s] %s %s", color.Red("PANIC RECOVER"), e, stack[:length])
	})

	if env.IsProduction {
		e.Use(middleware.Recover())
	}

	e.Use(middleware.Gzip())

	if env.IsProduction {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"https://ikenox.info"},
		}))
	} else {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}))
	}

	e.Use(useAppEngine)

	// TODO routingはどのレイヤの責務？
	pc := controller.NewCommentController()
	e.GET("/comment", pc.List)
	e.POST("/comment", pc.PostComment)
	e.DELETE("/comment/:id", pc.Delete)

	ec := controller.NewDomainEventController()
	e.POST("/_ah/push-handlers/domain-event", ec.Dispatch)
	return e
}

func useAppEngine(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if r, ok := c.Request().(*standard.Request); ok {
			ctx := appengine.WithContext(c.StdContext(), r.Request)
			ctx, err := appengine.Namespace(ctx, env.Namespace)
			if err != nil {
				panic(fmt.Sprintf("unresolve to set namespace (err %v)", err))
			}
			c.SetStdContext(ctx)
		}
		return next(c)
	}
}
