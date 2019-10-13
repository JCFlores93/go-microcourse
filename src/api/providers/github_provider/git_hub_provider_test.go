package github_provider

import (
	"errors"
	"github.com/JCFlores93/go-microcourse/src/api/clients/restclient"
	"github.com/JCFlores93/go-microcourse/src/api/domain/github"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestclient(t *testing.T) {
	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid resclient response"),
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)

	restclient.StopMockups()

	response, err = CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
}
