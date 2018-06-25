package repository

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type repository struct {
}

type dataStoreRepository struct {
	repository
	ctx  context.Context
	kind string
}

func newDataStoreRepository(ctx context.Context, kind string) *dataStoreRepository {
	return &dataStoreRepository{
		ctx:  ctx,
		kind: kind,
	}
}

func (r *dataStoreRepository) nextID() int64 {
	low, _, err := datastore.AllocateIDs(r.ctx, r.kind, nil, 1)
	if err != nil {
		panic(err.Error())
	}
	return low
}

// TODO: move to more concrete layer

func (r *dataStoreRepository) newKey(id int64) *datastore.Key {
	return datastore.NewKey(r.ctx, r.kind, "", id, nil)
}

func (r *dataStoreRepository) delete(key *datastore.Key) {
	err := datastore.Delete(r.ctx, key)
	if err != nil {
		panic(err.Error())
	}
}

func (r *dataStoreRepository) put(key *datastore.Key, src interface{}) *datastore.Key {
	key, err := datastore.Put(r.ctx, key, src)
	if err != nil {
		panic(err.Error())
	}
	return key
}

func (r *dataStoreRepository) get(key *datastore.Key, src interface{}) {
	err := datastore.Get(r.ctx, key, src)
	if err != nil {
		panic(err.Error())
	}
}

func (r *dataStoreRepository) newQuery() *datastore.Query {
	return datastore.NewQuery(r.kind)
}
