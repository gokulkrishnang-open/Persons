package middleware

import (
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

func TestAuthorizeUser(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST("/", AuthorizeUser)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", nil)
	c.Request.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJ1c2VyMiIsImV4cCI6MTY3NTc1NzM1MSwiaWF0IjoxNjc1NzUwMTUxfQ.jiAhOTq5JMrCm5pFvoiFPmycV4Lr9CY88TML5vwHW0I")
	r.ServeHTTP(w, c.Request)
	assert.NotEqual(t, 401, w.Code)

	w = httptest.NewRecorder()
	c.Request, _ = http.NewRequest(http.MethodPost, "/", nil)
	c.Request.Header.Add("Authorization", "Bearer bakdshbfjhad")
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 401, w.Code)
}
