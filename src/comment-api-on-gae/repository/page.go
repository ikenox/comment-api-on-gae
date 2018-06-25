package repository

import (
	"comment-api-on-gae/domain"
	"comment-api-on-gae/usecase"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type pageRepository struct {
	usecase.PageRepository
	*dataStoreRepository
}

func NewPageRepository(ctx context.Context) usecase.PageRepository {
	return &pageRepository{
		dataStoreRepository: newDataStoreRepository(ctx, "Page"),
	}
}

func (r *pageRepository) NextPageId() domain.PageId {
	return domain.PageId(r.nextID())
}

func (r *pageRepository) Add(page *domain.Page) {
	key, entity := r.toDataStoreEntity(page)
	r.put(key, entity)
}

func (r *pageRepository) Delete(id domain.PageId) {
	r.delete(r.newKey(int64(id)))
}

func (r *pageRepository) FindByUrl(pageUrl *domain.PageUrl) *domain.Page {
	return nil
}

type pageEntity struct {
	PageUrl string
}

func (r *pageRepository) toDataStoreEntity(page *domain.Page) (*datastore.Key, *pageEntity) {
	key := r.newKey(int64(page.PageId()))
	entity := &pageEntity{
		PageUrl: page.PageUrl().Url(),
	}
	return key, entity
}
