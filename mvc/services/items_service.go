package services

import (
	"github.com/bright-poku/go-microservice/mvc/domain"
	"github.com/bright-poku/go-microservice/mvc/utils"
	"net/http"
)

type itemsService struct {}

var (
	itemService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.AppError)  {
	return nil, &utils.AppError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "",
	}
}
