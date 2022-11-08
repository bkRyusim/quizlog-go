package controller

import (
	"github.com/bkRyusim/quizlog-go/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type PostController struct {
	TistoryService *service.TistoryService `inject:""`
}

func (p *PostController) GetAllPosts(ctx *fiber.Ctx) error {
	auth := ctx.Locals("user").(*jwt.Token)
	claims := auth.Claims.(jwt.MapClaims)
	accessToken := claims["access_token"].(string)

	posts, err := p.TistoryService.GetAllPosts(accessToken)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(posts)
}
