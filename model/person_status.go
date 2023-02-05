package model

type PersonAuthReq struct {
	UserName string `json:"DEFAULT_USERNAME"`
	Password string `json:"DEFAULT_PASSWORD"`
}

type PersonRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
}

type PersonResponse struct {
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
}
