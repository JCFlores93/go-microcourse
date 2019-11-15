package services

import (
	"github.com/JCFlores93/go-microcourse/src2/api/config"
	"github.com/JCFlores93/go-microcourse/src2/api/domain/github"
	"github.com/JCFlores93/go-microcourse/src2/api/domain/repositories"
	"github.com/JCFlores93/go-microcourse/src2/api/providers/github_provider"
	"github.com/JCFlores93/go-microcourse/src2/api/utils/errors"

)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request repositories.CreateRepoRequest) (repositories.CreateRepositoriesResult, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate();  err != nil {
		return nil, err
	}
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

func (s *reposService) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	ouput := make(chan repositories.CreateReposResponse)
	go s.handlerRepoResult(input, ouput)
	for _, current := range request {
		go s.CreateRepoConcurrent(current, input)
	}

}

func (s *reposService) handlerRepoResult(input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse
	for incommingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incommingEvent.Response,
			Error:    incommingEvent.Error,
		}

		results.Results = append(results.Results, repoResult)
	}
	output <- results
}

func (s *reposService) CreateRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) (repositories.CreateReposResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error:    err}
		return
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}
	result, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	output <- repositories.CreateRepositoriesResult{
		Response: &repositories.CreateRepoResponse{
			Id:    result.Id,
			Owner: result.Name,
			Name:  result.Owner.Login,
		},
		Error:    errors.NewApiError(err.StatusCode, err.Message),
	}
}