package models

type Users struct {
	ID       string
	Username string
}

func (t *Users) TableName() string {
	return "users"
}
