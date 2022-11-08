package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type HttpClient struct {
}

func New() *HttpClient {
	return &HttpClient{}
}

func (h *HttpClient) Get(url string, params url.Values) ([]byte, error) {
	requestUrl := url + "?" + params.Encode()
	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(requestUrl)
	fmt.Println(string(data))
	return data, nil
}
