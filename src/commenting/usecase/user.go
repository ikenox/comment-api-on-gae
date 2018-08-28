package usecase

import "commenting/domain/auth"

func NewUserUsecase(ur UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

type UserUsecase struct {
	userRepo UserRepository
}

func (u *UserUsecase) GetCurrentUser() (*auth.User, *Result) {
	return u.userRepo.CurrentUser(), &Result{
		code: OK,
	}
}
