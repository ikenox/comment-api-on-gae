package usecase

import (
	"domain"
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

func (c *CommentUseCase) Post(comment string) {
	post := &domain.Post{}
	c.postRepository.Add(post)
}
