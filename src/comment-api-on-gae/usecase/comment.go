package usecase

import (
	"comment-api-on-gae/domain"
)

// PostRepository stores Post
type PostRepository interface {
	Add(post *domain.Post)
	Delete(post *domain.Post)
	FindByPage(page *domain.Page)
}

type CommentUseCase struct {
	postRepository PostRepository
}

func NewCommentUseCase(repository PostRepository) *CommentUseCase {
	return &CommentUseCase{
		postRepository: repository,
	}
}

func (c *CommentUseCase) Post(pageUrl string, comment string) {
	post := &domain.Post{}
	c.postRepository.Add(post)
}
