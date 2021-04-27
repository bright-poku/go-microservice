package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bright-poku/go-microservice/mvc/services"
	"github.com/bright-poku/go-microservice/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		appErr := &utils.AppError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		jsonVal, _ := json.Marshal(appErr)
		resp.WriteHeader(appErr.StatusCode)
		resp.Write(jsonVal)
		return
	}

	user, appErr := services.GetUser(userId)
	if appErr != nil {
		jsonVal, _ := json.Marshal(appErr)
		resp.WriteHeader(appErr.StatusCode)
		resp.Write([]byte(jsonVal))
		return
	}
	//return content type of response
	jsonVal, _ := json.Marshal(user)
	resp.Write(jsonVal)
}
