package repository

import (
	"domain"
	"usecase"
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
