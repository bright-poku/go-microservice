package domain

import (
	"fmt"
	"github.com/bright-poku/go-microservice/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User {
		123: {
			Id:        123,
			FirstName: "bright",
			LastName:  "poku",
			Email:     "bap8000@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.AppError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.AppError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_founud",
	}
}
