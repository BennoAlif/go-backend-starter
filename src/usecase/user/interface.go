package userusecase

import (
	user "github.com/BennoAlif/go-backend-starter/src/repositories/user"
)

type sUserUsecase struct {
	userRepository user.UserRepository
}

type UserUsecase interface {
	CreateUser(*ParamsCreateUser) (*ResultLogin, error)
	Login(*ParamsLogin) (*ResultLogin, error)
}

func New(
	userRepository user.UserRepository,
) UserUsecase {
	return &sUserUsecase{
		userRepository: userRepository,
	}
}
