package repository

import (
	"context"
	"fmt"
	"github.com/bkRyusim/quizlog-go/domain"
	"github.com/bkRyusim/quizlog-go/ent"
	"github.com/bkRyusim/quizlog-go/ent/user"
)

type UserRepository struct {
	client *ent.UserClient
}

func NewUserRepository(client *ent.UserClient) *UserRepository {
	repository := UserRepository{}

	repository.client = client

	return &repository
}

func (u *UserRepository) Find(id int) (*domain.User, error) {
	user, err := u.client.Query().Where(user.ID(id)).Only(context.TODO())

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	result := domain.NewUser(user.ID, user.UserId, user.Name)

	return result, nil
}

func (u *UserRepository) Create(user domain.User) (int, error) {
	result, err := u.client.Create().SetName(user.Name()).SetUserId(user.UserId()).Save(context.TODO())
	if err != nil {
		return -1, fmt.Errorf("failed creating user: %w", err)
	}
	return result.ID, nil
}
