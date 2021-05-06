package services

import (
	"net/http"
	"testing"

	"github.com/bright-poku/go-microservice/mvc/domain"
	"github.com/bright-poku/go-microservice/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.AppError)
)

func init() {
	domain.UserData = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.AppError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.AppError) {
		return nil, &utils.AppError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 does not exists",
		}
	}

	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.AppError) {
		return &domain.User{
			Id: 123,
		}, nil
	}

	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
