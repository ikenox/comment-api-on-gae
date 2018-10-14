package presenter

import (
	"common/usecase"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type responseJson struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RenderJSON(c echo.Context, json interface{}, result *usecase.Result) error {
	var status int
	switch result.Code() {
	case usecase.OK:
		status = http.StatusOK
	case usecase.INVALID:
		status = http.StatusBadRequest
	case usecase.NOTFOUND:
		status = http.StatusNotFound
	case usecase.UNEXPECTED:
		status = http.StatusInternalServerError
	default:
		panic(fmt.Sprintf("Unknown Result Code '%s'", result.Code()))
	}

	return c.JSON(
		status,
		&responseJson{
			Message: result.Message(),
			Data:    json,
		},
	)
}

type jsonTime struct {
	time.Time
}

func (j jsonTime) format() string {
  return j.Time.UTC().Format("2006-01-02T15:04:05-07:00")
}

// MarshalJSON() の実装
func (j jsonTime) MarshalJSON() ([]byte, error) {
  return []byte(`"` + j.format() + `"`), nil
}
