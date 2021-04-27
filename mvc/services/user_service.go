package services

import "github.com/bright-poku/go-microservice/mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
