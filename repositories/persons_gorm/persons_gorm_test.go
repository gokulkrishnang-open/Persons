package person

import (
	db "persons/database"
	"persons/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	db.Connection()
}
func TestFetchEveryoneFromDB(t *testing.T) {
	page_args := model.Pagination{
		ResultsPerPage: "10",
		Page:           "1",
	}
	result, err := PersonGormImp{}.FetchEveryoneFromDB(page_args)
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
		UserName: "user25",
		Password: "password",
		Name:     "User Name",
		Email:    "user@gmail.com",
		Phone:    4921,
	}
	err := PersonGormImp{}.InsertPersonIntoDB(person)
	assert.Empty(t, err)
}

func TestDeletePersonFromDB(t *testing.T) {
	user := "user25"
	err := PersonGormImp{}.DeletePersonFromDB(user)
	assert.Empty(t, err)
}
