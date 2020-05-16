package services

import (
	"github.com/olegnikitashin/bookstore_users-api/domains/users"
	"github.com/olegnikitashin/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
