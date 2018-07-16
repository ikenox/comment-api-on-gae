package repository

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// TODO: ここらへんリファクタ
type dataStoreRepository struct {
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

func (r *dataStoreRepository) newKey(intId int64, stringId string) *datastore.Key {
	return datastore.NewKey(r.ctx, r.kind, stringId, intId, nil)
}

// エラーは基本的にpanicしている
// 必要に応じてリトライ機構とかつけてもよいが基本infra層のエラーはinfra層内で片付けたほうがほかの層がややこしくならない
// infra層のエラーがアプリ要件に絡んでくるようなら頑張ってerror返す？
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

func (r *dataStoreRepository) getMulti(keys []*datastore.Key, src interface{}) {
	if len(keys) == 0 {
		return
	}
	err := datastore.GetMulti(r.ctx, keys, src)
	if err == nil {
		return
	}
	mErr, ok := err.(appengine.MultiError)
	if !ok {
		panic(err.Error())
	}

	srcArr, ok := src.([]*struct{})
	for i, e := range mErr {
		if e == nil {
			continue
		}
		if e == datastore.ErrNoSuchEntity {
			srcArr[i] = nil
		}
		panic(e.Error())
	}
}

func (r *dataStoreRepository) get(key *datastore.Key, src interface{}) (ok bool) {
	err := datastore.Get(r.ctx, key, src)
	if err == nil {
		return true
	} else if err == datastore.ErrNoSuchEntity {
		return false
	} else {
		panic(err.Error())
	}
}

func (r *dataStoreRepository) newQuery() *datastore.Query {
	return datastore.NewQuery(r.kind)
}
