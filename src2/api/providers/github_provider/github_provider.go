package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/JCFlores93/go-microcourse/src2/api/clients/restclient"
	"github.com/JCFlores93/go-microcourse/src2/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	// Authorization "d02d8742f4a094a8a3d539dbee942d2b0a7a0c24"
	return fmt.Sprint(headerAuthorizationFormat, accessToken)
}
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	header := getAuthorizationHeader(accessToken)
	headers := http.Header{}
	headers.Set(headerAuthorization, header)
	response, err := restclient.Post(urlCreateRepo, headers, request)
	if err != nil {
		message := fmt.Sprintf("error when trying to create new repo in github %s", err.Error())
		log.Println(message)
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer response.Body.Close()
	if response.StatusCode > 299 {
		var errResponse = github.GithubErrorResponse{}
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}
	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		message := fmt.Sprintf("error when trying to create new repo successful response %s", err.Error())
		log.Println(message)
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error unmarshal"}
	}
	return &result, nil
}
