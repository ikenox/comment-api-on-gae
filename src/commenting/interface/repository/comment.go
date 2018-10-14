package repository

import (
	"commenting/domain"
	"commenting/usecase"
	"commenting/infra"
	"commenting/util"
	"context"
	"time"
)

type commentRepository struct {
	dao *infra.DataStoreDAO
}

func NewCommentRepository(ctx context.Context) usecase.CommentRepository {
	return &commentRepository{
		dao: infra.NewDataStoreDAO(ctx, "Comment"),
	}
}

func (r *commentRepository) NextCommentID() domain.CommentID {
	return domain.CommentID(r.dao.NextID())
}

func (r *commentRepository) Get(commentID domain.CommentID) *domain.Comment {
	entity := &commentEntity{
		_kind:r.dao.Kind(),
		CommentID: int64(commentID),
	}
	if ok := r.dao.Get(entity); ok {
		return r.fromDataStoreEntity(entity)
	} else {
		return nil
	}
}

func (r *commentRepository) Put(comment *domain.Comment) {
	entity := r.toDataStoreEntity(comment)
	r.dao.Put(entity)
}

func (r *commentRepository) Delete(commentId domain.CommentID) {
	r.dao.Delete(r.dao.NewKey("", int64(commentId), nil))
}

func (r *commentRepository) FindByPageID(pageId domain.PageID) []*domain.Comment {
	entities := []commentEntity{}
	query := r.dao.NewQuery().Filter("PageID =", pageId).Order("CommentedAt")
	_, err := r.dao.GetAll(query, &entities)
	if err != nil {
		panic(err.Error())
	}

	comments := make([]*domain.Comment, len(entities))
	for i, entity := range entities {
		comments[i] = r.fromDataStoreEntity(&entity)
	}
	return comments
}

type commentEntity struct {
	_kind       string `goon:"kind,U"`
	CommentID   int64  `datastore:"-" goon:"id"`
	PageID      string
	Text        []byte
	UserID      string
	Name        string
	CommentedAt time.Time
}

func (r *commentRepository) toDataStoreEntity(comment *domain.Comment) *commentEntity {
	return &commentEntity{
		_kind:       r.dao.Kind(),
		CommentID:   int64(comment.CommentID()),
		PageID:      string(comment.PageID()),
		Text:        util.StringToBytes(comment.Text()),
		UserID:      comment.Commenter().UserID(),
		Name:        comment.Name(),
		CommentedAt: comment.CommentedAt(),
	}
}

func (r *commentRepository) fromDataStoreEntity(entity *commentEntity) *domain.Comment {
	return domain.NewComment(
		domain.CommentID(entity.CommentID),
		domain.PageID(entity.PageID),
		util.BytesToString(entity.Text),
		entity.Name,
		domain.NewCommenter(
			entity.UserID,
		),
		entity.CommentedAt,
	)
}
