package repository

import (
	"commenting/domain"
	"commenting/usecase"
	"commenting/env"
	"context"
	"firebase.google.com/go/auth"
)

func NewCommenterRepository(ctx context.Context) usecase.CommenterRepository {
	client, err := env.FirebaseApp.Auth(ctx)
	if err != nil {
		panic(err.Error())
	}
	return &commenterRepository{
		authCli: client,
		ctx:     ctx,
	}
}

type commenterRepository struct {
	authCli *auth.Client
	ctx     context.Context
}

func (r *commenterRepository) CurrentCommenter(idToken string) *domain.Commenter {
	token, err := r.authCli.VerifyIDToken(r.ctx, idToken)
	if err != nil {
		return domain.NewCommenter("")
	}
	return domain.NewCommenter(token.UID)
}
