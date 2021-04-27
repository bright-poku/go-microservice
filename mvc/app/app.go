package app

import (
	"net/http"

	"github.com/bright-poku/go-microservice/mvc/controllers"
)

func StartApp()  {
	http.HandleFunc("/users", controllers.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}