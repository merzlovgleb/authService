package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itpark/market/auth/internal/domain"
	"github.com/itpark/market/auth/internal/infrastructure/repository"
	"github.com/itpark/market/auth/internal/presentation/http/user/dto"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &domain.User{
		ID:        uuid.New(),
		Name:      userDto.Name,
		Email:     userDto.Email,
		Password:  string(hashedPassword),
		Surname:   userDto.Surname,
		Role:      userDto.Role,
		CreatedAt: userDto.CreatedAt,
		IsActive:  userDto.IsActive,
	}

	_, err = userService.Repository.CreateUser(ctx, user)
	return err
}

func (userService *UserService) GetAllUsers(ctx *gin.Context) ([]domain.User, error) {
	return userService.Repository.GetAllUsers(ctx)
}

func (userService *UserService) GetUserById(ctx *gin.Context, id string) (*domain.User, error) {
	return userService.Repository.GetUserById(ctx, id)
}
