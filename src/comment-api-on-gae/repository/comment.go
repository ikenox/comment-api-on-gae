package repository

import (
	"comment-api-on-gae/domain"
	"comment-api-on-gae/usecase"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"time"
)

type commentRepository struct {
	*dataStoreRepository
	usecase.CommentRepository
}

type commentEntity struct {
	pageId      int64
	text        string
	commenterId int64
	commentedAt time.Time
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
	r.delete(r.newKey(int64(commentId)))
}

func (r *commentRepository) FindByPageId(pageId domain.PageId) []*domain.Comment {
	q := r.newQuery()
	var commentEntities []commentEntity
	keys, err := q.GetAll(r.ctx, &commentEntities)
	log.Infof(r.ctx, fmt.Sprint(len(keys)))
	log.Infof(r.ctx, fmt.Sprint(len(keys)))
	log.Infof(r.ctx, fmt.Sprint(len(keys)))
	log.Infof(r.ctx, fmt.Sprint(len(keys)))
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
	key := r.newKey(int64(comment.CommentId()))
	entity := &commentEntity{
		pageId:      int64(comment.PageId()),
		text:        comment.Text(),
		commenterId: int64(comment.CommenterId()),
		commentedAt: comment.CommentedAt(),
	}
	return key, entity
}

func (r *commentRepository) build(key *datastore.Key, entity *commentEntity) *domain.Comment {
	return domain.NewComment(
		domain.CommentId(key.IntID()),
		domain.PageId(entity.pageId),
		entity.text,
		domain.CommenterId(entity.commenterId),
		entity.commentedAt,
	)
}
