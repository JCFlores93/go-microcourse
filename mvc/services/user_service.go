package services

import "github.com/JCFlores93/go-microcourse/mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
