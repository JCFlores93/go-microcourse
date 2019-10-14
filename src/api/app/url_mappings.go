package app

import (
	"github.com/JCFlores93/go-microcourse/src/api/controllers/polo"
	"github.com/JCFlores93/go-microcourse/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
