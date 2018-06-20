package repository

import (
	"comment-api-on-gae/domain"
	"comment-api-on-gae/usecase"
	"golang.org/x/net/context"
)

type commenterRepository struct {
	*dataStoreRepository
	usecase.CommenterRepository
}

func NewCommenterRepository(ctx context.Context) usecase.CommenterRepository {
	return &commenterRepository{
		dataStoreRepository: newDataStoreRepository(ctx),
	}
}

func (c *commenterRepository) NextCommenterId() domain.CommenterId {
	return 1
}

func (c *commenterRepository) Add(post *domain.Commenter) {
}

func (c *commenterRepository) Delete(post *domain.Commenter) {
}

func (c *commenterRepository) FindById(page *domain.Page) *domain.Commenter {
	return domain.NewCommenter(111, "Bob")
}
