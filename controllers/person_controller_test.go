package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"persons/config"
	db "persons/database"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadConfigs()
	db.Connection()
}

func TestFindEveryoneHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/", FindEveryoneHandler)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}

func TestFindPersonHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/:username", FindPersonHandler)
	c.Request, _ = http.NewRequest(http.MethodGet, "/user1", nil)
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	c.Request, _ = http.NewRequest(http.MethodGet, "/user1@)#&$", nil)
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 400, w.Code)

	w = httptest.NewRecorder()
	c.Request, _ = http.NewRequest(http.MethodGet, "/user130818", nil)
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 500, w.Code)

}

func TestCreatePersonHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST("/", CreatePersonHandler)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`
	{
		"User_name": "user596*&(#@#$#",
		"Password": "password",
		"Name": "User Doe",
		"Email": "user@gmail.com",
		"Phone": 1234567891
	}
	`)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 400, w.Code)

	w = httptest.NewRecorder()
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`
	{
		"User_name": "user516",
		"Password": "password",
		"Name": "User Doe",
		"Email": "user@gmail.com",
		"Phone": 1234567891
	}
	`)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`
	{
		"User_name": "user516",
		"Password": "password",
		"Name": "User Doe",
		"Email": "user@gmail.com",
		"Phone": 1234567891
	}
	`)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 500, w.Code)

}

func TestUpdatePersonHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.PUT("/:username", UpdatePersonHandler)
	c.Request, _ = http.NewRequest(http.MethodPut, "/user516", bytes.NewBuffer([]byte(`
	{
		"User_name": "user516",
		"Password": "password",
		"Name": "User Name 516",
		"Email": "user@gmail.com",
		"Phone": 1234567891
	}
	`)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}

func TestDeletePersonHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.DELETE("/:username", DeletePersonHandler)
	c.Request, _ = http.NewRequest(http.MethodDelete, "/user516", nil)
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	c.Request, _ = http.NewRequest(http.MethodDelete, "/user535(*&$!&@", nil)
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 400, w.Code)
}
