package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	// Initialization:

	// Execution:
	user, err := GetUser(0)

	// Validation:
	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Jean", user.FirstName)
	assert.EqualValues(t, "Flores", user.LastName)
	assert.EqualValues(t, "jean@gmail.com", user.Email)
}