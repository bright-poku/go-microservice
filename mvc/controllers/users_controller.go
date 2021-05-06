package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/bright-poku/go-microservice/mvc/services"
	"github.com/bright-poku/go-microservice/mvc/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		appErr := &utils.AppError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		utils.RespondError(c, appErr)
		return // return response to the client
	}

	user, appErr := services.UsersService.GetUser(userId)
	if appErr != nil {
		utils.RespondError(c, appErr)
		return
	}
	//return content type of response
	utils.Respond(c, http.StatusOK, user)
}
