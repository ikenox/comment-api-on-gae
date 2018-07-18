package repository

import (
	"commenting/domain"
	"commenting/usecase"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"
	"util"
)

type commentRepository struct {
	*dataStoreRepository
}

type commentEntity struct {
	PageId      string
	Text        []byte
	CommenterId int64
	CommentedAt time.Time
}

func NewCommentRepository(ctx context.Context) usecase.CommentRepository {
	return &commentRepository{
		dataStoreRepository: newDataStoreRepository(ctx, "Comment"),
	}
}

func (r *commentRepository) NextCommentId() domain.CommentId {
	return domain.CommentId(r.nextID())
}

func (r *commentRepository) Add(comment *domain.Comment) {
	key, entity := r.toDataStoreEntity(comment)
	r.put(key, entity)
}

func (r *commentRepository) Delete(commentId domain.CommentId) {
	r.delete(r.newKey(int64(commentId), ""))
}

func (r *commentRepository) FindByPageId(pageId domain.PageId) []*domain.Comment {
	q := r.newQuery()
	var commentEntities []commentEntity
	keys, err := q.Filter("PageId =", pageId).GetAll(r.ctx, &commentEntities)
	if err != nil {
		panic(err.Error())
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
		PageId:      string(comment.PageId()),
		Text:        util.StringToBytes(comment.Text()),
		CommenterId: int64(comment.CommenterId()),
		CommentedAt: comment.CommentedAt(),
	}
	return key, entity
}

func (r *commentRepository) build(key *datastore.Key, entity *commentEntity) *domain.Comment {
	return domain.NewComment(
		domain.CommentId(key.IntID()),
		domain.PageId(entity.PageId),
		util.BytesToString(entity.Text),
		domain.CommenterId(entity.CommenterId),
		entity.CommentedAt,
	)
}
