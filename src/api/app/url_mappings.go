package app

import (
	"github.com/JCFlores93/go-microcourse/src/api/controllers/repositories"
	"github.com/JCFlores93/go-microcourse/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepo)
}