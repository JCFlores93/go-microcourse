package services

import (
	"github.com/JCFlores93/go-microcourse/src/api/config"
	"github.com/JCFlores93/go-microcourse/src/api/domain/github"
	"github.com/JCFlores93/go-microcourse/src/api/domain/repositories"
	"github.com/JCFlores93/go-microcourse/src/api/providers/github_provider"
	"github.com/JCFlores93/go-microcourse/src/api/utils/errors"
	"strings"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Private:     false,
		Description: input.Description,
	}
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Owner: response.Name,
		Name:  response.Owner.Login,
	}
	return &result, nil
}
