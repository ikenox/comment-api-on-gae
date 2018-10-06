package controller

import (
	u_notification "comment-api-on-gae/notification/usecase"
	"fmt"
	"github.com/labstack/echo"
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

	event := &struct {
		EventType string                 `json:"eventType"`
		Data      map[string]interface{} `json:"data"`
	}{}
	if err := c.Bind(event); err != nil {
		return err
	}
	println(fmt.Sprintf("%#v", event))

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
	}

	return nil
}
