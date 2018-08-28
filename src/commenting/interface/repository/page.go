package repository

import (
	"commenting/domain/comment"
	"commenting/usecase"
	"context"
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

func (r *pageRepository) Add(page *comment.Page) {
	key, entity := r.toDataStoreEntity(page)
	r.put(key, entity)
}

func (r *pageRepository) Delete(id comment.PageID) {
	r.delete(r.newKey(0, string(id)))
}

func (r *pageRepository) Get(id comment.PageID) *comment.Page {
	entity := new(pageEntity)
	key := r.newKey(0, string(id))
	ok := r.get(key, entity)
	if !ok {
		return nil
	}
	return r.build(key, entity)
}

type pageEntity struct{}

// TODO: repositoryが持ってんのなんか変
// presetnerに寄ってるのが正しい姿？
func (r *pageRepository) toDataStoreEntity(page *comment.Page) (*datastore.Key, *pageEntity) {
	key := r.newKey(0, string(page.PageId()))
	entity := &pageEntity{}
	return key, entity
}

func (r *pageRepository) build(key *datastore.Key, entity *pageEntity) *comment.Page {
	return comment.NewPage(comment.PageID(key.StringID()))
}
