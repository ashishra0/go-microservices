package services

import (
	"github.com/ashishra0/go-microservices/mvc/domain"
	"github.com/ashishra0/go-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}