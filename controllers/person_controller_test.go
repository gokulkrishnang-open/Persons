package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"persons/config"
	db "persons/database"
	"testing"
	"time"

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
	c.Request, _ = http.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte(``)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}

func TestFindPersonHandler(t *testing.T) {
	// username := "user1"
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.GET("/:username", FindPersonHandler)
	c.Request, _ = http.NewRequest(http.MethodGet, "/user1", bytes.NewBuffer([]byte(``)))
	r.ServeHTTP(w, c.Request)
	time.Sleep(300 * time.Millisecond)
	assert.Equal(t, 200, w.Code)
}

func TestCreatePersonHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST("/", CreatePersonHandler)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`
		"Username": "user50",
		"Password": "password",
		"Name": "User",
		"Email": "user50@gmail.com",
		"Phone": 1234567891,
	`)))
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)
}
