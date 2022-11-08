package repository

import (
	"context"
	"fmt"
	"github.com/bkRyusim/quizlog-go/domain"
	"github.com/bkRyusim/quizlog-go/ent"
	"github.com/bkRyusim/quizlog-go/ent/quiz"
	"github.com/bkRyusim/quizlog-go/ent/user"
	"github.com/bkRyusim/quizlog-go/request"
)

type QuizRepository struct {
	Client     *ent.QuizClient `inject:"quizClient"`
	UserClient *ent.UserClient `inject:"userClient"`
}

func (q *QuizRepository) Find(id int) (*domain.Quiz, error) {
	quiz, err := q.Client.Query().Where(quiz.ID(id)).Only(context.TODO())
	fmt.Println(quiz.Edges.User.UserId)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	result := domain.Quiz{Question: quiz.Question, Answer: quiz.Answer, PostUrl: quiz.PostUrl}
	return &result, nil
}

func (q *QuizRepository) Create(id string, quiz *request.Quiz) (int, error) {
	user, err := q.UserClient.Query().Where(user.UserId(id)).Only(context.TODO())
	result, err := q.Client.Create().SetQuestion(quiz.Question).SetAnswer(quiz.Answer).SetPostUrl(quiz.PostUrl).SetUser(user).Save(context.TODO())
	if err != nil {
		return -1, fmt.Errorf("failed creating quiz: %w", err)
	}
	return result.ID, nil
}

func (q *QuizRepository) FindAllByUserId(id string) ([]*ent.Quiz, error) {
	user, err := q.UserClient.Query().Where(user.UserId(id)).Only(context.TODO())
	all, err := user.QueryQuiz().All(context.TODO())
	return all, err
}
