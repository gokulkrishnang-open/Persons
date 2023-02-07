package controller

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// func init() {
// 	config.LoadConfigs()
// 	db.Connection()
// }

func TestLoginHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST("/", LoginHandler)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", nil)
	creds := base64.StdEncoding.EncodeToString([]byte("user1:password"))
	c.Request.Header.Add("Authorization", "Basic "+creds)
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	c.Request, _ = http.NewRequest(http.MethodPost, "/", nil)
	fakeCreds := base64.StdEncoding.EncodeToString([]byte("user1&#^@*:password"))
	c.Request.Header.Add("Authorization", "Basic "+fakeCreds)
	r.ServeHTTP(w, c.Request)
	assert.Equal(t, 401, w.Code)

}
