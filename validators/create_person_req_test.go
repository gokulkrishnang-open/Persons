package validator

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"persons/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidatePersonRequest(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{gin.Param{Key: "username", Value: "user123"}}

	res, err := ValidatePersonRequest(c)
	assert.NotEmpty(t, res)
	assert.Empty(t, err)
}

func TestValidateCreatePersonRequest(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	person := model.PersonRequest{
		UserName: "username",
		Password: "password",
		Name:     "User Name",
		Email:    "username@gmail.com",
		Phone:    1234567891,
	}
	persondata, _ := json.Marshal(&person)
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   ioutil.NopCloser(bytes.NewBuffer(persondata)),
	}
	res, err := ValidateCreatePersonRequest(c)
	assert.Empty(t, err)
	assert.NotEmpty(t, res)
}

func TestValidateUpdatePersonRequest(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	person := model.PersonRequest{
		UserName: "username",
		Password: "password",
		Name:     "User Name",
		Email:    "username@gmail.com",
		Phone:    1234567891,
	}
	persondata, _ := json.Marshal(&person)
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   ioutil.NopCloser(bytes.NewBuffer(persondata)),
	}
	res, err := ValidateUpdatePersonRequest(c)
	assert.Empty(t, err)
	assert.NotEmpty(t, res)
}
