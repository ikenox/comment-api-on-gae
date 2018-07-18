package controller

import (
	"commenting/interface/presenter"
	"commenting/usecase"
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
		status = http.StatusNotFound
	case usecase.E_NEXPECTED:
		status = http.StatusInternalServerError
	default:
		panic(fmt.Sprintf("Unknown Result Code '%s'", result.Code()))
	}

	return c.JSON(
		status,
		presenter.RenderJson(json, result),
	)
}
