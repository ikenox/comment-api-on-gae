package repository

import (
	"commenting/domain/comment"
	"commenting/usecase"
	"context"
	"google.golang.org/appengine/datastore"
)

type commenterRepository struct {
	*dataStoreRepository
}

type commenterEntity struct {
	Name string
}

func NewCommenterRepository(ctx context.Context) usecase.CommenterRepository {
	return &commenterRepository{
		dataStoreRepository: newDataStoreRepository(ctx, "Commenter"),
	}
}

func (r *commenterRepository) NextCommenterID() comment.CommenterID {
	return comment.CommenterID(r.nextID())
}

func (r *commenterRepository) Add(commenter *comment.Commenter) {
	key, entity := r.toDataStoreEntity(commenter)
	r.put(key, entity)
}

func (r *commenterRepository) Delete(id comment.CommenterID) {
	r.delete(r.newKey(int64(id), ""))
}

func (r *commenterRepository) Get(commenterId comment.CommenterID) *comment.Commenter {
	entity := new(commenterEntity)
	key := r.newKey(int64(commenterId), "")
	ok := r.get(key, entity)
	if !ok {
		return nil
	}
	return r.build(key, entity)
}

func (r *commenterRepository) FindByComments(commenterIds []comment.CommenterID) []*comment.Commenter {
	entities := make([]*commenterEntity, len(commenterIds))
	keys := make([]*datastore.Key, len(commenterIds))
	for i, id := range commenterIds {
		keys[i] = r.newKey(int64(id), "")
	}
	r.getMulti(keys, entities)

	commenters := make([]*comment.Commenter, len(commenterIds))
	for i, keys := range keys {
		commenters[i] = comment.NewCommenter(comment.CommenterID(keys.IntID()), entities[i].Name)
	}
	return commenters
}

func (r *commenterRepository) toDataStoreEntity(commenter *comment.Commenter) (*datastore.Key, *commenterEntity) {
	key := r.newKey(int64(commenter.CommenterId()), "")
	entity := &commenterEntity{
		Name: commenter.Name(),
	}
	return key, entity
}

func (r *commenterRepository) build(key *datastore.Key, entity *commenterEntity) *comment.Commenter {
	return comment.NewCommenter(comment.CommenterID(key.IntID()), entity.Name)
}
