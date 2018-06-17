package repository

import "golang.org/x/net/context"

type repository struct{
}

type dataStoreRepository struct {
	repository
	context context.Context
}

func newDataStoreRepository(ctx context.Context) *dataStoreRepository {
	return &dataStoreRepository{
		context: ctx,
	}
}
