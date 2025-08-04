package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/itpark/market/auth/internal/domain"
	"github.com/itpark/market/auth/internal/infrastructure/repository"
	"github.com/itpark/market/auth/internal/presentation/http/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrExistingUser       = errors.New("email already in use")
	ErrUserNotActive      = errors.New("user account is not active")
	ErrBlackListed        = errors.New("access token is blacklisted")
	ErrMinLengthPswd      = errors.New("password length must be between 6 and 128 characters")
)

var _ AuthUseCase = (*Auth)(nil)

type AuthUseCase interface {
	Register(email string, password string, name string, surname string) (userId string, err error)
	Login(email string, password string) (accessToken, refreshToken string, err error)
	RefreshToken(refreshToken string) (accessToken string, err error)
	ValidateToken(accessToken string) (valid bool, err error)
	Logout(accessToken string) (err error)
}
type Auth struct {
	userRepo     repository.UserRepository
	blacklist    repository.BlackListRepository
	tokenService token.JWTToken
}

func NewAuthUseCase(userRepo repository.UserRepository, blacklist repository.BlackListRepository, tokenSvc token.JWTToken) AuthUseCase {
	return &Auth{
		userRepo:     userRepo,
		blacklist:    blacklist,
		tokenService: tokenSvc,
	}
}

func (uc *Auth) Register(email, password, name, surname string) (string, error) {
	if len(password) < 6 || len(password) > 128 {
		return "", ErrMinLengthPswd
	}

	existingUser, err := uc.userRepo.GetUserByEmail(context.Background(), email)
	if err == nil && existingUser != nil {
		return "", ErrExistingUser
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	userID := uuid.New()
	newUser := &domain.User{
		ID:        userID,
		Email:     email,
		Password:  string(hashedPassword),
		Name:      name,
		Surname:   surname,
		CreatedAt: time.Now(),
		IsActive:  true,
	}

	_, err = uc.userRepo.CreateUser(context.Background(), newUser)
	if err != nil {
		return "", err
	}

	return newUser.ID.String(), nil
}

func (uc *Auth) Login(email string, password string) (string, string, error) {
	curUser, err := uc.userRepo.GetUserByEmail(context.Background(), email)
	if err != nil || curUser == nil {
		return "", "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(curUser.Password), []byte(password)); err != nil {
		return "", "", ErrInvalidCredentials
	}

	if !curUser.IsActive {
		return "", "", ErrUserNotActive
	}
	accesToken, err := uc.tokenService.GenerateAccessToken(curUser)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := uc.tokenService.GenerateRefreshToken(curUser)
	if err != nil {
		return "", "", err
	}
	return accesToken, refreshToken, nil
}

func (uc *Auth) RefreshToken(refreshToken string) (string, error) {
	isValid, err := uc.checkToken(refreshToken)
	if !isValid || err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	newAccessToken, err := uc.tokenService.RefreshAccessToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %w", err)
	}
	return newAccessToken, nil
}

func (uc *Auth) ValidateToken(accessToken string) (bool, error) {
	return uc.checkToken(accessToken)
}

func (uc *Auth) checkToken(token string) (bool, error) {

	isBlacklisted, err := uc.blacklist.IsTokenBlacklisted(context.Background(), token)
	if err != nil {
		return false, fmt.Errorf("failed to check blacklist: %w", err)
	}
	if isBlacklisted {
		return false, ErrBlackListed
	}

	isValid, err := uc.tokenService.ValidateToken(token)
	if err != nil {
		return false, err
	}
	return isValid, nil
}

func (uc *Auth) Logout(accessToken string) error {
	err := uc.blacklist.AddToBlacklist(context.Background(), accessToken)
	if err != nil {
		return fmt.Errorf("failed to add token to blacklist: %w", err)
	}
	return nil
}
