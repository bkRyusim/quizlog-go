package service

import (
	"github.com/bkRyusim/quizlog-go/domain"
	"github.com/bkRyusim/quizlog-go/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository `inject:""`
}

func (u *UserService) Create(user *domain.User) (int, error) {
	id, err := u.UserRepository.Create(user)
	return id, err
}

func (u *UserService) Find(id int) (*domain.User, error) {
	user, err := u.UserRepository.Find(id)
	return user, err
}
