package person

import (
	"persons/model"
)

func GetPeople(raw_people []model.Persons) []model.PersonResponse {
	resp := make([]model.PersonResponse, len(raw_people))
	for i, raw_person := range raw_people {
		resp[i] = model.PersonResponse{
			UserName: raw_person.UserName,
			Name:     raw_person.Name,
			Email:    raw_person.Email,
			Phone:    raw_person.Phone,
		}
	}
	return resp
}

func GetPerson(raw_person model.Persons) model.PersonResponse {
	resp := model.PersonResponse{
		UserName: raw_person.UserName,
		Name:     raw_person.Name,
		Email:    raw_person.Email,
		Phone:    raw_person.Phone,
	}
	return resp
}

func GetPersonEntity(raw_person model.PersonRequest) model.Persons {
	resp := model.Persons{
		UserName: raw_person.UserName,
		Password: raw_person.Password,
		Name:     raw_person.Name,
		Email:    raw_person.Email,
		Phone:    raw_person.Phone,
	}
	return resp
}

func GetPersonAuthReq(raw_person model.Persons) model.PersonAuthReq {
	resp := model.PersonAuthReq{
		UserName: raw_person.UserName,
		Password: raw_person.Password,
	}
	return resp
}
