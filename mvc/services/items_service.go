package services

import (
	"github.com/JCFlores93/go-microcourse/mvc/domain"
	"github.com/JCFlores93/go-microcourse/mvc/utils"
	"net/http"
)

type itemsService struct {}

func (s *itemsService)GetIemt(itemId string) (*domain.Item, *utils.ApplicationError){
	return nil, &utils.ApplicationError{
		Message: "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}