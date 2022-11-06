package tistory

import (
	"encoding/json"
	client "github.com/bkRyusim/quizlog-go/client/http"
	"github.com/bkRyusim/quizlog-go/client/tistory/constants"
	"github.com/bkRyusim/quizlog-go/client/tistory/response"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type TistoryClient struct {
	accessToken  string
	clientId     string
	clientSecret string
	redirectUri  string
	code         string
	httpClient   *client.HttpClient
}

func New() *TistoryClient {
	t := TistoryClient{}
	t.clientId = os.Getenv("TISTORY_CLIENT_ID")
	t.clientSecret = os.Getenv("TISTORY_CLIENT_SECRET")
	t.redirectUri = os.Getenv("TISTORY_REDIRECT_URI")
	t.httpClient = client.New()

	return &t
}

// Init Initialize client
func (t *TistoryClient) Init(code string) error {
	params := url.Values{
		"client_id":     {t.clientId},
		"client_secret": {t.clientSecret},
		"redirect_uri":  {t.redirectUri},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	}

	data, err := t.httpClient.Get(constants.TISTORY_AUTH_URL, params)
	if err != nil {
		return err
	}

	t.accessToken = strings.Split(string(data), "=")[1]
	return nil
}

// GetBlogInfo Get blog info
func (t *TistoryClient) GetBlogInfo() (*response.TistoryBlogInfo, error) {
	params := url.Values{
		"access_token": {t.accessToken},
		"output":       {"json"},
	}

	data, err := t.httpClient.Get(constants.TISTORY_BLOG_INFO_URL, params)
	if err != nil {
		return nil, err
	}

	result := &response.TistoryBlogInfo{}

	json.Unmarshal(data, result)

	return result, nil
}

func (t *TistoryClient) GetPosts(blogName string, pageNumber int) (*response.TistoryPostList, error) {
	params := url.Values{
		"access_token": {t.accessToken},
		"output":       {"json"},
		"blogName":     {blogName},
		"page":         {strconv.Itoa(pageNumber)},
	}

	data, err := t.httpClient.Get(constants.TISTORY_BLOG_POST_LIST, params)
	if err != nil {
		return nil, err
	}

	result := &response.TistoryPostList{}

	json.Unmarshal(data, result)

	return result, nil
}
