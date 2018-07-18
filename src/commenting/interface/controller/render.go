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
	case usecase.ErrInvalid:
		status = http.StatusBadRequest
	case usecase.ErrNotFound:
		status = http.StatusNotFound
	case usecase.ErrUnexpected:
		status = http.StatusInternalServerError
	default:
		panic(fmt.Sprintf("Unknown Result Code '%s'", result.Code()))
	}

	return c.JSON(
		status,
		presenter.RenderJson(json, result),
	)
}
