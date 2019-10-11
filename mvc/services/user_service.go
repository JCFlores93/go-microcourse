package services

import (
	"github.com/JCFlores93/go-microcourse/mvc/domain"
	"github.com/JCFlores93/go-microcourse/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
