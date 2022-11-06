package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	assert := assert2.New(t)

	app := Setup()
	request := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(request)

	statusCode := resp.StatusCode
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(200, statusCode, "Status code mismatched")
	assert.Equal("Quizlog API", string(body), "Body mismatched")
}
