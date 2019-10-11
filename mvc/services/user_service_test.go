package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (userDaoMock userDaoMock)
type userDaoMock struct {}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)

}
