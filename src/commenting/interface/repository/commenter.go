package repository

import (
	"commenting/domain"
	"commenting/usecase"
	"golang.org/x/net/context"
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

func (r *commenterRepository) NextCommenterId() domain.CommenterId {
	return domain.CommenterId(r.nextID())
}

func (r *commenterRepository) Add(commenter *domain.Commenter) {
	key, entity := r.toDataStoreEntity(commenter)
	r.put(key, entity)
}

func (r *commenterRepository) Delete(id domain.CommenterId) {
	r.delete(r.newKey(int64(id), ""))
}

func (r *commenterRepository) Get(commenterId domain.CommenterId) *domain.Commenter {
	entity := new(commenterEntity)
	key := r.newKey(int64(commenterId), "")
	ok := r.get(key, entity)
	if !ok {
		return nil
	}
	return r.build(key, entity)
}

func (r *commenterRepository) FindByComments(commenterIds []domain.CommenterId) []*domain.Commenter {
	entities := make([]*commenterEntity, len(commenterIds))
	keys := make([]*datastore.Key, len(commenterIds))
	for i, id := range commenterIds {
		keys[i] = r.newKey(int64(id), "")
	}
	r.getMulti(keys, entities)

	commenters := make([]*domain.Commenter, len(commenterIds))
	for i, keys := range keys {
		commenters[i] = domain.NewCommenter(domain.CommenterId(keys.IntID()), entities[i].Name)
	}
	return commenters
}

func (r *commenterRepository) toDataStoreEntity(commenter *domain.Commenter) (*datastore.Key, *commenterEntity) {
	key := r.newKey(int64(commenter.CommenterId()), "")
	entity := &commenterEntity{
		Name: commenter.Name(),
	}
	return key, entity
}

func (r *commenterRepository) build(key *datastore.Key, entity *commenterEntity) *domain.Commenter {
	return domain.NewCommenter(domain.CommenterId(key.IntID()), entity.Name)
}
