package service

import (
	"fmt"
	"github.com/bkRyusim/quizlog-go/client/tistory"
	tistoryresponse "github.com/bkRyusim/quizlog-go/client/tistory/response"
	"github.com/bkRyusim/quizlog-go/config"
	"github.com/bkRyusim/quizlog-go/repository"
	"github.com/bkRyusim/quizlog-go/request"
	"github.com/bkRyusim/quizlog-go/response"
	"github.com/golang-jwt/jwt"
	"sync"
	"time"
)

type TistoryService struct {
	UserRepository *repository.UserRepository `inject:""`
}

func (t *TistoryService) Auth(request *request.Auth) (*response.Login, error) {
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

	generatedToken, err := token.SignedString([]byte(config.Config.AuthConfig.SecretKey))
	if err != nil {
		return nil, err
	}

	user, err := t.UserRepository.FindByUserId(userId)

	isNewUser := true
	name := ""

	if err == nil {
		isNewUser = false
		name = user.Name()
	}

	return &response.Login{Token: generatedToken, UserId: userId, IsNewUser: isNewUser, Name: name}, nil
}

func (t *TistoryService) GetAllPosts(token string) ([]tistoryresponse.Post, error) {
	tistoryClient := tistory.NewTistoryClient()
	tistoryClient.SetAccessToken(token)

	data, err := tistoryClient.GetBlogInfo()
	if err != nil {
		panic(err)
	}

	blogs := data.Tistory.Item.Blogs
	totalPost := 0
	for _, b := range blogs {
		totalPost += b.Statistics.Post/10 + 1
	}

	var wait sync.WaitGroup
	wait.Add(totalPost/10 + 1)

	type MutexPosts struct {
		Posts []tistoryresponse.Post
		sync.RWMutex
	}
	mu := MutexPosts{Posts: []tistoryresponse.Post{}}

	for _, blog := range blogs {
		for page := 1; page <= blog.Statistics.Post/10+1; page++ {
			go func(p int) {
				mu.Lock()
				defer mu.Unlock()
				defer wait.Done()
				fmt.Println(p)
				post, err := tistoryClient.GetPosts(blog.Name, p)
				if err != nil {
					panic(err)
				}
				mu.Posts = append(mu.Posts, post.Tistory.Item.Posts...)
			}(page)
		}
	}
	wait.Wait()

	return mu.Posts, nil
}
