package oauth

import (
	"github.com/JCFlores93/go-microcourse/outh-api/src/api/domain/oauth"
	"github.com/JCFlores93/go-microcourse/outh-api/src/api/services"
	"github.com/JCFlores93/go-microcourse/src/api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context) {
	token, err := services.OauthService.GetAccessToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, token)
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
}
