package model

type User struct {
	ID uint 			`json:"id"`
	Username string 	`json:"username"`
	Password string 	`json:"password"`
	Email string 		`json:"email"`
	PhoneNumber string 	`json:"phone number"`
}

type Lost struct {
	ID uint 			`json:"id"`
	Description string 	`json:"description"`
	UserID string 	`json:"user id"`
}

type Found struct {
	ID uint 			`json:"id"`
	Description string 	`json:"description"`
	UserID string 	`json:"user id"`
}