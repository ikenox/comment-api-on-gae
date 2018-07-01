package repository

import (
	"comment-api-on-gae/domain"
	"comment-api-on-gae/usecase"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type pageRepository struct {
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
	r.delete(r.newKey(0, string(id)))
}

func (r *pageRepository) Get(id domain.PageId) *domain.Page {
	entity := &pageEntity{}
	err := r.get(r.newKey(0, string(id)), entity)
	// TODO: ここまでdatastoreのerr引き回してくるのはあんまりきれいじゃない？
	if err == datastore.ErrNoSuchEntity {
		return nil
	}
	return domain.NewPage(id)
}

type pageEntity struct{}

func (r *pageRepository) toDataStoreEntity(page *domain.Page) (*datastore.Key, *pageEntity) {
	key := r.newKey(0, string(page.PageId()))
	entity := &pageEntity{}
	return key, entity
}
