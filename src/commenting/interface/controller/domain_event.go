package controller

import (
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
	//ctx := c.StdContext()

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
