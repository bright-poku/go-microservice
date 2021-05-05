package domain

import (
	"fmt"
	"net/http"

	"github.com/bright-poku/go-microservice/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "john",
			LastName:  "doe",
			Email:     "johndoe@gmail.com",
		},
		125: {
			Id:        125,
			FirstName: "mary",
			LastName:  "doe",
			Email:     "marydoe@gmail.com",
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
		Code:       "not_found",
	}
}
