package repository

import (
	"comment-api-on-gae/commenting/domain"
	"comment-api-on-gae/commenting/usecase"
	"comment-api-on-gae/common/infra"
	"context"
	"google.golang.org/appengine/datastore"
)

type pageRepository struct {
	dao *infra.DataStoreDAO
}

func NewPageRepository(ctx context.Context) usecase.PageRepository {
	return &pageRepository{
		dao: infra.NewDataStoreDAO(ctx, "Page"),
	}
}

func (r *pageRepository) Add(page *domain.Page) {
	key, entity := r.toDataStoreEntity(page)
	r.dao.Put(key, entity)
}

func (r *pageRepository) Delete(id domain.PageID) {
	r.dao.Delete(r.dao.NewKey(0, string(id)))
}

func (r *pageRepository) Get(id domain.PageID) *domain.Page {
	entity := new(pageEntity)
	key := r.dao.NewKey(0, string(id))
	ok := r.dao.Get(key, entity)
	if !ok {
		return nil
	}
	return r.build(key, entity)
}

type pageEntity struct{}

// TODO: repositoryが持ってんのなんか変
// presetnerに寄ってるのが正しい姿？
func (r *pageRepository) toDataStoreEntity(page *domain.Page) (*datastore.Key, *pageEntity) {
	key := r.dao.NewKey(0, string(page.PageId()))
	entity := &pageEntity{}
	return key, entity
}

func (r *pageRepository) build(key *datastore.Key, entity *pageEntity) *domain.Page {
	return domain.NewPage(domain.PageID(key.StringID()))
}
