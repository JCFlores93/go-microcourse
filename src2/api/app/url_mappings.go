package app

import (
	"github.com/JCFlores93/go-microcourse/src2/api/controllers/polo"
	"github.com/JCFlores93/go-microcourse/src2/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
