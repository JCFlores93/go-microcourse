package services

import (
	"github.com/JCFlores93/go-microcourse/mvc/domain"
	"github.com/JCFlores93/go-microcourse/mvc/utils"
)

type usersService struct {

}

var (
	UsersService usersService
)

func (u *usersService)GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
