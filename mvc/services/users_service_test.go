package services

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNotFoundInDatabase(t *testing.T) {
	user, err := UsersService.GetUser(1)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 1 was not found", err.Message)
}