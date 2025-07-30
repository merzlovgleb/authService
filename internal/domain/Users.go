package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	Name     string    `json:"name" db:"name"`
	Surname  string    `json:"surname" db:"surname"`
}

func NewUser(id uuid.UUID, email string, password string, name string, surname string) User {
	return User{
		ID:       id,
		Email:    email,
		Password: password,
		Name:     name,
		Surname:  surname,
	}
}
