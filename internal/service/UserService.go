package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/auth/internal/domain"
	"github.com/itpark/market/auth/internal/infrastructure/repository"
	"github.com/itpark/market/auth/internal/presentation/http/user/dto"
)

type UserService struct {
	Repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		Repository: repository,
	}
}

func (userService *UserService) CreateUser(ctx context.Context, userDto *dto.CreateUserDto) error {
	user := &domain.User{
		Name:      userDto.Name,
		Email:     userDto.Email,
		Password:  userDto.Password, //захэшировать надо бы
		Surname:   userDto.Surname,
		Role:      userDto.Role,
		CreatedAt: userDto.CreatedAt,
		IsActive:  userDto.IsActive,
	}

	_, err := userService.Repository.CreateUser(ctx, user)
	return err
}

func (userService *UserService) GetAllUsers(ctx *gin.Context) []domain.User {
	return userService.Repository.GetAllUsers(ctx)
}
