package repository

import (
	"comment-api-on-gae/domain"
	"comment-api-on-gae/usecase"
)

type PostDataStore struct {
	usecase.PostRepository
}

func (repo *PostDataStore) Add(post domain.Post) {

}
func (repo *PostDataStore) Delete(post domain.Post) {

}
func (repo *PostDataStore) FindByPage(page domain.Page) {

}
