package app

import (
	"github.com/JCFlores93/go-microcourse/outh-api/src/api/controllers/oauth"
	"github.com/JCFlores93/go-microcourse/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token")
}