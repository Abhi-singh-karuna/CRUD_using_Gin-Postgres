package model

type Employee struct {
	ID uint `gorm:"primary_key"`

	Names    string `json:"names"`
	Username string `json:"username"`
	Email    string ` json:"email"`
	Password string ` json:"password"`
}
