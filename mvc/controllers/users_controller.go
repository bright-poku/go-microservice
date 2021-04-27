package controllers

import (
	"encoding/json"
	"github.com/bright-poku/go-microservice/mvc/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(err.Error()))
		return
	}
	//return content type of response
	jsonVal, _ := json.Marshal(user)
	resp.Write(jsonVal)
}
