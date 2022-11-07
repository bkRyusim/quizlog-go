package controller

import (
	"github.com/bkRyusim/quizlog-go/request"
	"github.com/bkRyusim/quizlog-go/service"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	TistoryService *service.TistoryService `inject:""`
}

func (a *AuthController) Login(ctx *fiber.Ctx) error {
	authRequest := new(request.Auth)

	if err := ctx.BodyParser(authRequest); err != nil {
		return err
	}

	token, err := a.TistoryService.Auth(authRequest)
	if err != nil {
		return ctx.SendStatus(500)
	}

	return ctx.JSON(token)
}
