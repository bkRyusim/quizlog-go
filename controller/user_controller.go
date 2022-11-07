package controller

import (
	"fmt"
	"github.com/bkRyusim/quizlog-go/domain"
	"github.com/bkRyusim/quizlog-go/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserController struct {
	UserService *service.UserService `inject:""`
}

func (u *UserController) NewUser(ctx *fiber.Ctx) error {
	ctx.Body()
	user := domain.NewUser(0, "123456", "백승윤")
	id, err := u.UserService.Create(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(strconv.Itoa(id))
	return ctx.SendString(strconv.Itoa(id))
}
