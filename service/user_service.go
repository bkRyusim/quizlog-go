package service

import "github.com/bkRyusim/quizlog-go/repository"

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	service := UserService{}
	service.userRepository = userRepository
	return service
}
