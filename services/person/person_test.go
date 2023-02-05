package person

import (
	"persons/config"
	db "persons/database"
	"persons/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	config.LoadConfigs()
	db.Connection()
}

func TestFindEveryone(t *testing.T) {
	resp, err := PersonsImplementation{}.FindEveryone("1", "10")
	assert.Empty(t, err)
	assert.NotEmpty(t, resp)
}

func TestFindPerson(t *testing.T) {
	resp, err := PersonsImplementation{}.FindPerson("user1")
	assert.Empty(t, err)
	assert.NotEmpty(t, resp)
}

func TestCreatePerson(t *testing.T) {
	person_req := model.PersonRequest{
		UserName: "username32",
		Password: "password",
		Name:     "User 30",
		Email:    "user30@gmail.com",
		Phone:    1234567891,
	}
	err := PersonsImplementation{}.CreatePerson(person_req)
	assert.Empty(t, err)
}

func TestDeletePerson(t *testing.T) {
	err := PersonsImplementation{}.DeletePerson("username32")
	assert.Empty(t, err)
}

func TestFetchPersonCreds(t *testing.T) {
	resp, err := PersonsImplementation{}.FetchPersonCreds("user1")
	assert.Empty(t, err)
	assert.NotEmpty(t, resp)
}
