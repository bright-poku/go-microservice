package services

import (
	"github.com/bright-poku/go-microservice/mvc/domain"
	"github.com/bright-poku/go-microservice/mvc/utils"
)

type userService struct {

}

var (
	usersService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.AppError) {
	return domain.GetUser(userId)
}
