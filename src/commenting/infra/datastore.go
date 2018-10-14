package infra

import (
	"context"
	"github.com/mjibson/goon"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// TODO: ここらへんリファクタ
// middleware層行き？
type DataStoreDAO struct {
	ctx  context.Context
	kind string
	g    *goon.Goon
}

func (d *DataStoreDAO) Kind() string {
	return d.kind
}

func NewDataStoreDAO(ctx context.Context, kind string) *DataStoreDAO {
	return &DataStoreDAO{
		ctx:  ctx,
		g:    goon.FromContext(ctx),
		kind: kind,
	}
}
func (d *DataStoreDAO) NewKey(stringID string, intID int64, parent *datastore.Key) *datastore.Key {
	return datastore.NewKey(d.ctx, d.kind, stringID, intID, parent)
}

func (d *DataStoreDAO) NextID() int64 {
	low, _, err := datastore.AllocateIDs(d.ctx, d.kind, nil, 1)
	if err != nil {
		panic(err.Error())
	}
	return low
}

// TODO エラー時にリトライなどさせるべきか
func (d *DataStoreDAO) Delete(key *datastore.Key) {
	err := d.g.Delete(key)
	if err != nil {
		panic(err.Error())
	}
}

func (d *DataStoreDAO) Put(src interface{}) *datastore.Key {
	key, err := d.g.Put(src)
	if err != nil {
		panic(err.Error())
	}
	return key
}

func (d *DataStoreDAO) GetMulti(keys []*datastore.Key, src interface{}) {
	if len(keys) == 0 {
		return
	}
	err := datastore.GetMulti(d.ctx, keys, src)
	if err == nil {
		return
	}
	mErr, ok := err.(appengine.MultiError)
	if !ok {
		panic(err.Error())
	}

	for _, e := range mErr {
		if e == nil {
			continue
		}
		if e == datastore.ErrNoSuchEntity {
			continue
		}
		panic(e.Error())
	}
}

func (d *DataStoreDAO) GetAll(q *datastore.Query, list interface{}) ([]*datastore.Key, error) {
	return d.g.GetAll(q, list)
}

func (d *DataStoreDAO) Get(src interface{}) (ok bool) {
	err := d.g.Get(src)
	if err == nil {
		return true
	} else if err == datastore.ErrNoSuchEntity {
		return false
	} else {
		panic(err.Error())
	}
}

func (d *DataStoreDAO) NewQuery() *datastore.Query {
	return datastore.NewQuery(d.kind)
}
