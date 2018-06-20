package repository

import (
	"comment-api-on-gae/usecase"
	"comment-api-on-gae/domain"
	"golang.org/x/net/context"
)

type commentRepository struct {
	*dataStoreRepository
	usecase.CommentRepository
}

func NewCommentRepository(ctx context.Context) usecase.CommentRepository {
	return &commentRepository{
		dataStoreRepository: newDataStoreRepository(ctx),
	}
}

func (r *commentRepository) NextCommentId() domain.CommentId {
	return 1111
}

func (r *commentRepository) Add(post *domain.Comment) {
}

func (r *commentRepository) Delete(post *domain.Comment) {
}

func (r *commentRepository) FindByPage(page *domain.Page) {

}
