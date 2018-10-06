package repository

import (
	"cloud.google.com/go/pubsub"
	"comment-api-on-gae/env"
	"context"
	"encoding/json"
)

type EventPublisher struct {
	ctx context.Context
}

func NewPublisher(ctx context.Context) *EventPublisher {
	return &EventPublisher{ctx}
}

func (p *EventPublisher) Publish(eventType string, data interface{}) {
	PubsubClient, err := pubsub.NewClient(p.ctx, env.ProjectID, env.GCPCredentialOption)
	if err != nil {
		panic(err.Error())
	}

	t := PubsubClient.Topic("domain-event")
	if err != nil {
		panic(err.Error())
	}

	bytes, err := json.Marshal(struct {
		Data interface{} `json:"data"`
		EventType string `json:"eventType"`
	}{
		EventType: eventType,
		Data:      data,
	})
	if err != nil {
		panic(err.Error())
	}

	t.Publish(p.ctx, &pubsub.Message{Data: bytes})
	//_, err = result.Get(p.ctx)
	//if err != nil {
	//	panic(err.Error())
	//}
}
