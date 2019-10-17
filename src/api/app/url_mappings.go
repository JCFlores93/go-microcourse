package app

import (
	"github.com/JCFlores93/go-microcourse/src/api/controllers/polo"
	"github.com/JCFlores93/go-microcourse/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositores", repositories.CreateRepos)
}
