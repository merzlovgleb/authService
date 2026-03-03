package dto

import (
	_ "github.com/google/uuid"
	"github.com/itpark/market/auth/internal/domain"
)

type UpdateUserDto struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Name     *string `json:"name,omitempty"`
	Surname  *string `json:"surname,omitempty"`
	Role     *string `json:"role,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
}

// Apply применяет изменения из DTO к доменной модели
func (d *UpdateUserDto) Apply(user *domain.User) {
	if d.Email != nil {
		user.Email = *d.Email
	}
	if d.Password != nil {
		user.Password = *d.Password
	}
	if d.Name != nil {
		user.Name = *d.Name
	}
	if d.Surname != nil {
		user.Surname = *d.Surname
	}
	if d.Role != nil {
		user.Role = *d.Role
	}
	if d.IsActive != nil {
		user.IsActive = *d.IsActive
	}
}
