package app

import (
	"github.com/JCFlores93/go-microcourse/src/api/log/option_a"
	"github.com/gin-gonic/gin"
)

var( router *gin.Engine)

func init() {
	router = gin.Default()
}

func StartApp() {
	option_a.Log.Info("about to map the urls")
	mapUrls()
	option_a.Log.Info("urls successfully mapped")
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}