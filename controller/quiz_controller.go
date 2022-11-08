package controller

import (
	"github.com/bkRyusim/quizlog-go/request"
	"github.com/bkRyusim/quizlog-go/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type QuizController struct {
	QuizService *service.QuizService `inject:""`
}

func (q *QuizController) NewQuiz(ctx *fiber.Ctx) error {
	auth := ctx.Locals("user").(*jwt.Token)
	claims := auth.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)

	request := new(request.Quiz)
	ctx.BodyParser(request)

	q.QuizService.Create(userId, request)

	return ctx.SendStatus(200)
}

func (q *QuizController) GetAllQuiz(ctx *fiber.Ctx) error {
	auth := ctx.Locals("user").(*jwt.Token)
	claims := auth.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)

	res, err := q.QuizService.FindAll(userId)
	if err != nil {
		panic(err)
	}
	return ctx.JSON(res)
}
