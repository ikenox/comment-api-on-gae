package repository

import (
	"comment-api-on-gae/commenting/domain"
	"comment-api-on-gae/commenting/usecase"
	"comment-api-on-gae/common/infra"
	"comment-api-on-gae/env"
	"context"
	"firebase.google.com/go/auth"
	"google.golang.org/appengine/datastore"
)

func NewCommenterRepository(ctx context.Context) usecase.CommenterRepository {
	client, err := env.FirebaseApp.Auth(ctx)
	if err != nil {
		panic(err.Error())
	}
	return &commenterRepository{
		authCli: client,
		dao:     infra.NewDataStoreDAO(ctx, "Commenter"),
		ctx:     ctx,
	}
}

type commenterRepository struct {
	authCli *auth.Client
	dao     *infra.DataStoreDAO
	ctx     context.Context
}

func (r *commenterRepository) NextCommenterID() domain.CommenterID {
	return domain.CommenterID(r.dao.NextID())
}

func (r *commenterRepository) Put(commenter *domain.Commenter) {
	key, entity := r.toDataStoreEntity(commenter)
	r.dao.Put(key, entity)
}

func (r *commenterRepository) FindByCommenterID(commenterIDs []domain.CommenterID) []*domain.Commenter {
	entities := make([]*commenterEntity, len(commenterIDs))
	keys := make([]*datastore.Key, len(commenterIDs))
	for i, id := range commenterIDs {
		keys[i] = r.dao.NewKey(int64(id), "")
	}
	r.dao.GetMulti(keys, entities)

	commenters := make([]*domain.Commenter, len(commenterIDs))
	for i, key := range keys {
		if entities[i] != nil {
			commenters[i] = r.build(key, entities[i])
		} else {
			commenters[i] = nil
		}
	}
	return commenters
}

func (r *commenterRepository) CurrentCommenter(idToken string) *domain.Commenter {
	token, err := r.authCli.VerifyIDToken(r.ctx, idToken)
	if err != nil {
		return nil
	}
	userID := domain.UserID(token.UID)

	// todo repositoryがロジック持ってる
	// userIDの扱いがなにかおかしいかも
	commenter := r.getByUserID(userID)
	if commenter == nil {
		commenter = domain.NewCommenter(r.NextCommenterID(), "", userID)
		r.Put(commenter)
	}
	return commenter
}

func (r *commenterRepository) getByUserID(userID domain.UserID) *domain.Commenter {
	var entities []commenterEntity
	keys, err := r.dao.NewQuery().Filter("UserID =", string(userID)).Limit(1).GetAll(r.ctx, &entities)
	if err != nil {
		panic(err.Error())
	}
	if len(keys) == 0 {
		return nil
	}

	key := keys[0]
	entity := entities[0]
	return r.build(key, &entity)
}

type commenterEntity struct {
	Name   string
	UserID string
}

func (r *commenterRepository) toDataStoreEntity(commenter *domain.Commenter) (*datastore.Key, *commenterEntity) {
	key := r.dao.NewKey(int64(commenter.CommenterID()), "")
	entity := &commenterEntity{
		Name:   commenter.Name(),
		UserID: string(commenter.UserID()),
	}
	return key, entity
}

func (r *commenterRepository) build(key *datastore.Key, entity *commenterEntity) *domain.Commenter {
	return domain.NewCommenter(domain.CommenterID(key.IntID()), entity.Name, domain.UserID(entity.UserID))
}
