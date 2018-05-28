package repository

import (
	"comment-api-on-gae/domain"
	"comment-api-on-gae/usecase"
	"context"
)

type postDataStore struct {
	usecase.PostRepository
	context context.Context
}

func NewPostRepository(ctx context.Context) *postDataStore {
	return &postDataStore{
		context: ctx,
	}
}

func (repo *postDataStore) Add(post *domain.Post) {

}

func (repo *postDataStore) Delete(post *domain.Post) {

}

func (repo *postDataStore) FindByPage(page *domain.Page) {

}
