package repository

import (
	"comment-api-on-gae/commenting/domain"
	"comment-api-on-gae/commenting/usecase"
	"comment-api-on-gae/common/infra"
	"comment-api-on-gae/util"
	"context"
	"google.golang.org/appengine/datastore"
	"time"
)

type commentRepository struct {
	dao *infra.DataStoreDAO
	ctx context.Context
}

func NewCommentRepository(ctx context.Context) usecase.CommentRepository {
	return &commentRepository{
		dao: infra.NewDataStoreDAO(ctx, "Comment"),
		ctx: ctx,
	}
}

func (r *commentRepository) NextCommentID() domain.CommentID {
	return domain.CommentID(r.dao.NextID())
}

func (r *commentRepository) Put(comment *domain.Comment) {
	key, entity := r.toDataStoreEntity(comment)
	r.dao.Put(key, entity)
}

func (r *commentRepository) Delete(commentId domain.CommentID) {
	r.dao.Delete(r.dao.NewKey(int64(commentId), ""))
}

func (r *commentRepository) FindByPageID(pageId domain.PageID) []*domain.Comment {
	q := r.dao.NewQuery()
	var commentEntities []commentEntity
	keys, fuga := q.Filter("PageID =", pageId).Order("CommentedAt").GetAll(r.ctx, &commentEntities)
	if fuga != nil {
		panic(fuga.Error())
	}

	comments := make([]*domain.Comment, len(keys))
	for i, key := range keys {
		comments[i] = r.build(key, &commentEntities[i])
	}
	return comments
}

type commentEntity struct {
	PageID      string
	Text        []byte
	CommenterID int64
	CommentedAt time.Time
}

func (r *commentRepository) toDataStoreEntity(comment *domain.Comment) (*datastore.Key, *commentEntity) {
	key := r.dao.NewKey(int64(comment.CommentID()), "")
	entity := &commentEntity{
		PageID:      string(comment.PageID()),
		Text:        util.StringToBytes(comment.Text()),
		CommenterID: int64(comment.CommenterID()),
		CommentedAt: comment.CommentedAt(),
	}
	return key, entity
}

func (r *commentRepository) build(key *datastore.Key, entity *commentEntity) *domain.Comment {
	return domain.NewComment(
		domain.CommentID(key.IntID()),
		domain.PageID(entity.PageID),
		util.BytesToString(entity.Text),
		domain.CommenterID(entity.CommenterID),
		entity.CommentedAt,
	)
}
