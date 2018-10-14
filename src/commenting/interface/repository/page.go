package repository

import (
	"commenting/domain"
	"commenting/usecase"
	"common/infra"
	"context"
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
	r.dao.Put(&pageEntity{
		_kind:  r.dao.Kind(),
		PageID: string(page.PageID()),
	})
}

func (r *pageRepository) Delete(id domain.PageID) {
	r.dao.Delete(r.dao.NewKey(string(id), 0, nil))
}

func (r *pageRepository) Get(id domain.PageID) *domain.Page {
	entity := &pageEntity{
		_kind: r.dao.Kind(),
		PageID:string(id),
	}
	ok := r.dao.Get(entity)
	if !ok {
		return nil
	}
	return r.build(entity)
}

type pageEntity struct {
	_kind  string `goon:"kind,U"`
	PageID string `datastore:"-" goon:"id"`
}

func (r *pageRepository) build(entity *pageEntity) *domain.Page {
	return domain.NewPage(domain.PageID(entity.PageID))
}
