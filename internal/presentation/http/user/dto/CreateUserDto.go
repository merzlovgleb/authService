package dto

import "time"

type CreateUserDto struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Surname   string    `json:"surname"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}
