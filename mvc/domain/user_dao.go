package domain

import (
	"fmt"
	"github.com/ashishra0/go-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User {
		123: {Id: 123, FirstName: "Ashish", LastName: "Rao", Email: "ashish@gmail.com"},
	}

	UserDao userDao
)

type userDao struct {}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code: "not found",
	}
}