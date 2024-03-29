package controllers

import (
	"github.com/JCFlores93/go-microcourse/mvc/services"
	"github.com/JCFlores93/go-microcourse/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
	return
}

func respond(accept string) {

}
