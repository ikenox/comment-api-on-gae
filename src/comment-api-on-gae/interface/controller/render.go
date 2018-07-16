package controller

import (
	"comment-api-on-gae/interface/presenter"
	"comment-api-on-gae/usecase"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func renderJSON(c echo.Context, json interface{}, result *usecase.Result) error {
	var status int
	switch result.Code() {
	case usecase.OK:
		status = http.StatusOK
	case usecase.E_INVALID:
		status = http.StatusBadRequest
	case usecase.E_NOTFOUND:
		status = http.StatusInternalServerError
	case usecase.E_NEXPECTED:
		status = http.StatusNotFound
	default:
		panic(fmt.Sprintf("Unknown Result Code '%s'", result.Code()))
	}

	return c.JSON(
		status,
		presenter.RenderJson(json, result),
	)
}
