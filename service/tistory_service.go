package service

import (
	"github.com/bkRyusim/quizlog-go/client/tistory"
	"github.com/bkRyusim/quizlog-go/request"
	"github.com/bkRyusim/quizlog-go/response"
	"github.com/golang-jwt/jwt"
	"time"
)

type TistoryService struct {
}

func (t *TistoryService) Auth(request *request.Auth) (*response.Token, error) {
	tistoryClient := tistory.NewTistoryClient()

	err := tistoryClient.Init(request.Code)
	if err != nil {
		return nil, err
	}

	result, err := tistoryClient.GetBlogInfo()
	if err != nil {
		return nil, err
	}
	userId := result.Tistory.Item.UserId
	accessToken := tistoryClient.AccessToken()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["access_token"] = accessToken
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	generatedToken, err := token.SignedString([]byte{})
	if err != nil {
		return nil, err
	}

	return &response.Token{Token: generatedToken}, nil
}
