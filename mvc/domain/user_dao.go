package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User {
		123: {
			Id:        1,
			FirstName: "bright",
			LastName:  "poku",
			Email:     "bap8000@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, error) {
	user := users[userId]
	if user == nil {
		return nil, errors.New(fmt.Sprintf("user %v was not found", userId))
	}
	return user, nil
}
