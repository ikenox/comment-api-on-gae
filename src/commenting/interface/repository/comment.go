package repository

import (
	"commenting/domain"
	"commenting/usecase"
	"context"
	"google.golang.org/appengine/datastore"
	"time"
	"util"
)

type commentRepository struct {
	*dataStoreRepository
}

type commentEntity struct {
	PageID      string
	Text        []byte
	CommenterID int64
	CommentedAt time.Time
}

func NewCommentRepository(ctx context.Context) usecase.CommentRepository {
	return &commentRepository{
		dataStoreRepository: newDataStoreRepository(ctx, "Comment"),
	}
}

func (r *commentRepository) NextCommentID() domain.CommentID {
	return domain.CommentID(r.nextID())
}

func (r *commentRepository) Add(comment *domain.Comment) {
	key, entity := r.toDataStoreEntity(comment)
	r.put(key, entity)
}

func (r *commentRepository) Delete(commentId domain.CommentID) {
	r.delete(r.newKey(int64(commentId), ""))
}

func (r *commentRepository) FindByPageID(pageId domain.PageID) []*domain.Comment {
	q := r.newQuery()
	var commentEntities []commentEntity
	keys, fuga := q.Filter("PageID =", pageId).GetAll(r.ctx, &commentEntities)
	if fuga != nil {
		panic(fuga.Error())
	}

	comments := make([]*domain.Comment, len(keys))
	for i, key := range keys {
		comments[i] = r.build(key, &commentEntities[i])
	}
	return comments
}

func (r *commentRepository) toDataStoreEntity(comment *domain.Comment) (*datastore.Key, *commentEntity) {
	key := r.newKey(int64(comment.CommentId()), "")
	entity := &commentEntity{
		PageID:      string(comment.PageId()),
		Text:        util.StringToBytes(comment.Text()),
		CommenterID: int64(comment.CommenterId()),
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
