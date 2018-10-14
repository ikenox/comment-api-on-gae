package controller

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"notification/usecase"
	"time"
)

type DomainEventController struct{}

func NewDomainEventController() *DomainEventController {
	return &DomainEventController{}
}

func (ctl *DomainEventController) Dispatch(c echo.Context) error {
	ctx := c.StdContext()

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
		pageId, _ := event.Data["pageId"].(string)
		name, _ := event.Data["name"].(string)
		text, _ := event.Data["text"].(string)
		usecase.NotifyCommentPosted(ctx, int64(commentId), pageId, name, text)
	case "CommentDeleted":
		// FIXME int64として受け取りたい
		commentId, _ := event.Data["commentId"].(float64)
		pageId, _ := event.Data["pageId"].(string)
		name, _ := event.Data["name"].(string)
		text, _ := event.Data["text"].(string)
		usecase.NotifyCommentDeleted(ctx, int64(commentId), pageId, name, text)
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
