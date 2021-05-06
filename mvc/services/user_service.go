package services

import (
	"github.com/bright-poku/go-microservice/mvc/domain"
	"github.com/bright-poku/go-microservice/mvc/utils"
)

type userService struct {}

var (
	UsersService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.AppError) {
	user, err :=  domain.UserData.GetUser(userId)
	if err != nil {
		return nil, err
	}

	return user, nil

}
