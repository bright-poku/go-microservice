package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit test: TestGetUserNotFound
// test with go test -cover
// test stages : initialization, execution, validation
func TestGetUserNotFound(t *testing.T) {

	// execution
	user, err := GetUser(0)

	// validation
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)

}

// test for return user
// expecting the user exists
func TestGetUser(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "john", user.FirstName)
	assert.EqualValues(t, "doe", user.LastName)
}