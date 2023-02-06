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
func TestFetchEveryoneFromDB(t *testing.T) {
	pageArgs := model.Pagination{
		ResultsPerPage: "10",
		Page:           "1",
	}
	result, err := PersonGormImp{}.FetchEveryoneFromDB(pageArgs)
	assert.Empty(t, err)
	assert.NotEmpty(t, result)
}

func TestFetchPersonFromDB(t *testing.T) {
	user := "user1"
	res, err := PersonGormImp{}.FetchPersonFromDB(user)
	assert.Empty(t, err)
	assert.NotEmpty(t, res)
}

func TestInsertPersonIntoDB(t *testing.T) {
	person := model.Persons{
		UserName: "user27",
		Password: "password",
		Name:     "User Name",
		Email:    "user@gmail.com",
		Phone:    4921,
	}
	err := PersonGormImp{}.InsertPersonIntoDB(person)
	assert.Empty(t, err)
}

func TestUpdatePersonInDB(t *testing.T) {
	userName := "user27"
	person := model.Persons{
		UserName: "user27",
		Password: "password",
		Name:     "User Name27",
		Email:    "user@gmail.com",
		Phone:    4921,
	}
	err := PersonGormImp{}.UpdatePersonInDB(userName, person)
	assert.Empty(t, err)
}

func TestDeletePersonFromDB(t *testing.T) {
	user := "user27"
	err := PersonGormImp{}.DeletePersonFromDB(user)
	assert.Empty(t, err)
}
