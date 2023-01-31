package validator

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"persons/model"
)

func ValidatePersonRequest(ctx *gin.Context) (string, error) {
	type req struct {
		UserName string `json:"user_name"`
	}
	reqbody := req{ctx.Params.ByName("username")}
	rules := govalidator.MapData{
		"user_name": []string{"regex:^[a-zA-Z0-9]+$"},
	}
	messages := govalidator.MapData{
		"user_name": []string{"regex:Invalid Username"},
	}
	opts := govalidator.Options{
		Rules:    rules,
		Messages: messages,
		Data:     &reqbody,
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		err := map[string]interface{}{"error": e}
		ctx.JSON(400, err)
		ctx.Abort()
		return reqbody.UserName, errors.New("Invaid request")
	}
	return reqbody.UserName, nil

}

func ValidateCreatePersonRequest(ctx *gin.Context) (model.PersonRequest, error) {
	person := model.PersonRequest{}
	ctx.BindJSON(&person)

	rules := govalidator.MapData{
		"user_name": []string{"required", "min:5", "max:100", "regex:^[a-zA-Z0-9]+$"},
		"password":  []string{"required", "min:8", "max:15"},
		"name":      []string{"required", "min:1", "max:100"},
		"email":     []string{"required", "min:4", "max:20", "email"},
		"phone":     []string{"required", "digits:10", "regex:^[0-9]+$"},
	}
	messages := govalidator.MapData{
		"user_name": []string{"required:Username is required!!!", "min:Username should be atleast 5 character long", "max:Username should not exceed 100 characters", "regex:Invalid Character in Username"},
		"password":  []string{"required:Password is required!!", "min:password should be atleast 8 characters long", "max:Password should not exceed 15 characters"},
		"name":      []string{"required:Name is required!!", "min:Name should be atleast 1 character long", "max:Name should not exceed 100 characters"},
		"email":     []string{"required:Email is required!!", "min:Email should be atleast 4 characters long", "max:Email should not exceed 20 characters"},
		"phone":     []string{"required:Phone number is required", "digits:Phone number should be 10 digits", "regex:Please enter a valid phone number"},
	}
	opts := govalidator.Options{
		Rules:    rules,
		Messages: messages,
		Data:     &person,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		err := map[string]interface{}{"error": e}
		ctx.JSON(400, err)
		ctx.Abort()
		return person, errors.New("Invalid request")
	}

	return person, nil
}

func ValidateUpdatePersonRequest(ctx *gin.Context) (model.PersonRequest, error) {
	person := model.PersonRequest{}
	ctx.BindJSON(&person)
	fmt.Println(person)

	rules := govalidator.MapData{
		"user_name": []string{"min:5", "max:100", "regex:^[a-zA-Z0-9]+$"},
		"password":  []string{"min:8", "max:15"},
		"name":      []string{"min:1", "max:100"},
		"email":     []string{"min:4", "max:20", "email"},
		"phone":     []string{"digits:10", "regex:^[0-9]+$"},
	}
	messages := govalidator.MapData{
		"user_name": []string{"min:Username should be atleast 5 character long", "max:Username should not exceed 100 characters", "regex:Invalid Character in Username"},
		"password":  []string{"min:password should e atleast 8 characters long", "max:Password should not exceed 15 characters"},
		"name":      []string{"min:Name should be atleast 1 character long", "max:Name should not exceed 100 characters"},
		"email":     []string{"min:Email should be atleast 4 characters long", "max:Email should not exceed 20 characters"},
		"phone":     []string{"digits:Phone number should be 10 digits", "regex:Please enter a valid phone number"},
	}
	opts := govalidator.Options{
		Rules:    rules,
		Messages: messages,
		Data:     &person,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		ctx.JSON(400, gin.H{"error": e})
		ctx.Abort()
		return person, errors.New("Invalid request")
	}

	return person, nil
}
