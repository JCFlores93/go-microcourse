package app

import (
	"github.com/JCFlores93/go-microcourse/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
