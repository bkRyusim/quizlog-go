package controller

import (
	"github.com/bkRyusim/quizlog-go/domain"
	"github.com/bkRyusim/quizlog-go/request"
	"github.com/bkRyusim/quizlog-go/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

type UserController struct {
	UserService *service.UserService `inject:""`
}

func (u *UserController) NewUser(ctx *fiber.Ctx) error {
	joinRequest := new(request.Join)
	if err := ctx.BodyParser(joinRequest); err != nil {
		return err
	}

	auth := ctx.Locals("user").(*jwt.Token)
	claims := auth.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)

	user := domain.NewUser(0, userId, joinRequest.Name)
	id, err := u.UserService.Create(user)
	if err != nil {
		panic(err)
	}
	return ctx.SendString(strconv.Itoa(id))
}
