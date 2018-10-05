package repository

import (
	"cloud.google.com/go/pubsub"
	"comment-api-on-gae/env"
	"context"
)


type EventPublisher struct {
	ctx context.Context
}

func NewPublisher(ctx context.Context) *EventPublisher {
	return &EventPublisher{ctx}
}

func (p *EventPublisher) Publish(topicID string, message string) {
	PubsubClient, err := pubsub.NewClient(p.ctx, env.ProjectId, env.GCPCredentialOption)
	if err != nil {
		panic(err.Error())
	}

	PubsubClient.Topic(topicID)
	t := PubsubClient.Topic(topicID)
	if err != nil {
		panic(err.Error())
	}

	result := t.Publish(p.ctx, &pubsub.Message{Data: []byte( message )})
	_, err = result.Get(p.ctx)
	if err != nil {
		panic(err.Error())
	}
}
