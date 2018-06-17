package repository

import (
	"comment-api-on-gae/domain"
	"comment-api-on-gae/usecase"
	"golang.org/x/net/context"
)

type pageRepository struct {
	usecase.PageRepository
	*dataStoreRepository
}

func NewPageRepository(ctx context.Context) usecase.PageRepository {
	return &pageRepository{
		dataStoreRepository: newDataStoreRepository(ctx),
	}
}

func (r *pageRepository) NextPageId() *domain.PageId {
	panic("implement me")
}

func (r *pageRepository) Add(post *domain.Comment) {
	panic("implement me")
}

func (r *pageRepository) Delete(post *domain.Comment) {
	panic("implement me")
}

func (r *pageRepository) FindByUrl(url string) *domain.Page {
	panic("implement me")
}
