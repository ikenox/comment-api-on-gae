package repository

import (
	"commenting/env"
	"golang.org/x/net/context"
)

type FirebaseUserRepository struct {
	ctx context.Context
}

func (r *FirebaseUserRepository) GetUser() {
	env.FirebaseApp.Auth(r.ctx)
}

func NewFirebaseUserRepository(ctx context.Context) (*FirebaseUserRepository, error) {
	return &FirebaseUserRepository{
		ctx: ctx,
	}, nil
}
