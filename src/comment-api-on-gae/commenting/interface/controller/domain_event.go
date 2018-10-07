package controller

import (
	u_notification "comment-api-on-gae/notification/usecase"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type DomainEventController struct{}

func NewDomainEventController() *DomainEventController {
	return &DomainEventController{}
}

func (ctl *DomainEventController) Dispatch(c echo.Context) error {
	ctx := c.StdContext()

	// TODO subscribeに関する実装のリファクタ、クラスの適切な分割など
	// TODO 各コンテキストに1つずつdispatcherを用意してそれらをhandlerから呼び出すようにする
	// TODO イベントタイプは意図しない重複防ぐためにenumかなんかで定義、stringでハードコードしないように改修
	msg := &pushRequest{}
	if err := json.NewDecoder(c.Request().Body()).Decode(msg); err != nil {
		return c.HTML(http.StatusBadRequest, fmt.Sprintf("Could not decode body: %v", err))
	}
	event := struct {
		Data      map[string]interface{} `json:"data"`
		EventType string                 `json:"eventType"`
	}{}
	err := json.Unmarshal([]byte(msg.Message.Data), &event)
	if err != nil {
		return c.HTML(http.StatusBadRequest, fmt.Sprintf("Could not unmarshal message: %v", err))
	}

	switch event.EventType {
	case "CommentPosted":
		// FIXME int64として受け取りたい
		commentId, _ := event.Data["commentId"].(float64)
		name, _ := event.Data["name"].(string)
		text, _ := event.Data["text"].(string)
		u_notification.NotifyCommentPosted(ctx, int64(commentId), name, text)
	case "CommentDeleted":
		// FIXME int64として受け取りたい
		commentId, _ := event.Data["commentId"].(float64)
		name, _ := event.Data["name"].(string)
		text, _ := event.Data["text"].(string)
		u_notification.NotifyCommentDeleted(ctx, int64(commentId), name, text)
	default:
		return c.JSON(http.StatusBadRequest, nil)
	}

	return nil
}

type pushRequest struct {
	Message struct {
		Data        []byte    `json:"data"`
		MessageID   string    `json:"messageId"`
		PublishTime time.Time `json:"publishTime"`
	}
	Subscription string
}
