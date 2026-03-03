package service

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	return &UserService{Repository: repository}
}

func (s *UserService) CreateUser(ctx context.Context, userDto *dto.CreateUserDto) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	role := userDto.Role
	if role == "" {
		role = "user"
	}

	user := &domain.User{
		ID:        uuid.New(),
		Name:      userDto.Name,
		Surname:   userDto.Surname,
		Email:     userDto.Email,
		Password:  string(hashedPassword),
		Role:      role,
		IsActive:  true,
		CreatedAt: time.Now().UTC(),
	}

	createdUser, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser.Password = ""
	return createdUser, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) []domain.User {
	return s.Repository.GetAllUsers(ctx)
}

func (s *UserService) GetUserById(ctx context.Context, id string) (*domain.User, error) {
	return s.Repository.GetUserById(ctx, id)
}
func (s *UserService) DeleteUserById(ctx context.Context, id string) error {
	return s.Repository.DeleteUserById(ctx, id)
}

func (s *UserService) UpdateUserById(ctx context.Context, id string, dto *dto.UpdateUserDto) (*domain.User, error) {
	// 1. Получаем текущего пользователя
	existingUser, err := s.Repository.GetUserById(ctx, id)
	if err != nil {
		return nil, err // пробрасываем sql.ErrNoRows, если не найден
	}

	// 2. Применяем изменения из DTO
	dto.Apply(existingUser)

	// 3. (Опционально) Валидация бизнес-логики
	if existingUser.Email == "" {
		return nil, errors.New("email cannot be empty")
	}

	// 4. Хэшируем пароль, если он изменён
	if dto.Password != nil && *dto.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}
		existingUser.Password = string(hashedPassword)
	}

	// 5. Сохраняем изменения
	updatedUser, err := s.Repository.UpdateUserById(ctx, existingUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user in repository: %w", err)
	}

	return updatedUser, nil
}
