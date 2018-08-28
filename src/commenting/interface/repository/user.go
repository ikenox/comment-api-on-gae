package repository

import (
	"commenting/domain/auth"
	"commenting/usecase"
	"context"
	"google.golang.org/appengine/user"
)

type appEngineUserRepository struct {
	ctx context.Context
}

func NewUserRepository(ctx context.Context) usecase.UserRepository {
	return &appEngineUserRepository{
		ctx: ctx,
	}
}

func (r *appEngineUserRepository) CurrentUser() *auth.User {
	if u := user.Current(r.ctx); u != nil {
		return auth.NewUser(u.ID)
	} else {
		return nil
	}
}
