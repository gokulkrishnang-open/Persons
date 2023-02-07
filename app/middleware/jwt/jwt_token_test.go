package jwt

import (
	"persons/config"
	"persons/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadConfigs()
}

func TestCreateToken(t *testing.T) {
	req := model.PersonAuthReq{
		UserName: "user1",
		Password: "password",
	}
	resp, err := CreateToken(req)
	assert.NotEmpty(t, resp)
	assert.Nil(t, err)
}

func TestValidateJwtAuthToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJ1c2VyMiIsImV4cCI6MTY3NTc1NzM1MSwiaWF0IjoxNjc1NzUwMTUxfQ.jiAhOTq5JMrCm5pFvoiFPmycV4Lr9CY88TML5vwHW0I"
	err := ValidateJwtAuthToken(token)
	assert.Nil(t, err)
	fakeToken := "lhfadlskjfh"
	err = ValidateJwtAuthToken(fakeToken)
	assert.NotNil(t, err)
}
