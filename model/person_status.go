package model

type PersonRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type PersonResponse struct {
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
}
