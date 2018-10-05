package repository

import (
	"comment-api-on-gae/commenting/usecase"
	"context"
	"google.golang.org/appengine/log"
)

type loggingRepository struct {
	ctx context.Context
}

func (r *loggingRepository) Infof(format string, args ...interface{}){
	log.Infof(r.ctx, format, args)
}

func NewLoggingRepository(ctx context.Context) usecase.LoggingRepository{
	return &loggingRepository{ctx}
}
