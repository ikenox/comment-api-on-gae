package repository

import (
	"commenting/env"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/pubsub/v1"
)

type EventPublisher struct {
	ctx context.Context
}

func NewPublisher(ctx context.Context) *EventPublisher {
	return &EventPublisher{ctx}
}

func (p *EventPublisher) Publish(eventType string, data interface{}) {
	bytes, err := json.Marshal(struct {
		Data      interface{} `json:"data"`
		EventType string      `json:"eventType"`
	}{
		EventType: eventType,
		Data:      data,
	})
	if err != nil {
		panic(err.Error())
	}

	hcli, err := google.DefaultClient(p.ctx, pubsub.PubsubScope)
	if err != nil {
		panic(err.Error())
	}

	pubsubService, err := pubsub.New(hcli)
	if err != nil {
		panic(err.Error())
	}

	// TODO 開発環境だとレスポンス遅いがGCP環境だと大丈夫かどうか
	_, err = pubsubService.Projects.Topics.Publish(
		fmt.Sprintf("projects/%s/topics/domain-event", env.ProjectID),
		&pubsub.PublishRequest{
			Messages: []*pubsub.PubsubMessage{
				{
					Data: base64.StdEncoding.EncodeToString(bytes),
				},
			},
		},
	).Do()
	if err != nil {
		panic(err.Error())
	}
}
