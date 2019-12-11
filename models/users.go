package models

/*
Users Users

*/
type Users struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string             `json:"password"`
	Score string `json:"score"`
	Email string `json:"email"`
	LastLoginDate string `json:"lastLoginDate"`
}
