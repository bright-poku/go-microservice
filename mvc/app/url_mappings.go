package app

import (
	"github.com/bright-poku/go-microservice/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
