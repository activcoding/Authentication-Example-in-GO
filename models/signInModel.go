package models

type SignInModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
