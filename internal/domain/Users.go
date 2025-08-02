package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Name      string    `json:"name" db:"name"`
	Surname   string    `json:"surname" db:"surname"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}

func NewUser(id uuid.UUID, email string, password string, name string, surname string, role string, createdAt time.Time, isActive bool) User {
	return User{
		ID:        id,
		Email:     email,
		Password:  password,
		Name:      name,
		Surname:   surname,
		Role:      role,
		CreatedAt: createdAt,
		IsActive:  isActive,
	}
}
