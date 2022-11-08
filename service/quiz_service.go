package service

import (
	"fmt"
	"github.com/bkRyusim/quizlog-go/ent"
	"github.com/bkRyusim/quizlog-go/repository"
	"github.com/bkRyusim/quizlog-go/request"
)

type QuizService struct {
	QuizRepository *repository.QuizRepository `inject:""`
}

func (q *QuizService) Create(id string, request *request.Quiz) (int, error) {
	res, err := q.QuizRepository.Create(id, request)
	fmt.Println(res)
	return res, err
}

func (q *QuizService) FindAll(id string) ([]*ent.Quiz, error) {
	res, err := q.QuizRepository.FindAllByUserId(id)
	return res, err
}
