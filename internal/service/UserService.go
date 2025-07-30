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

func (userService *UserService) CreateUser(ctx context.Context, userDto dto.CreateUserDto) {
	userService.Repository.CreateUser(ctx, userDto.Name, userDto.Email, userDto.Password)
}

func (userService *UserService) GetAllUsers(ctx *gin.Context) []domain.User {
	return userService.Repository.GetAllUsers(ctx)
}
