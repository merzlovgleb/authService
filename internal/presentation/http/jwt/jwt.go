package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/itpark/market/auth/internal/domain"
	"time"
)

var _ JWTToken = (*jwtToken)(nil)

var (
	ErrMissingSecret       = errors.New("missing JWT_SECRET in config")
	ErrInvalidToken        = errors.New("invalid JWT token")
	ErrAccessTokenExpired  = errors.New("access token expired")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
)

type JWTToken interface {
	GenerateAccessToken(user *domain.User) (string, error)
	GenerateRefreshToken(user *domain.User) (string, error)
	ValidateToken(token string) (bool, error)
	RefreshAccessToken(refreshToken string) (string, error)
}

type jwtToken struct {
	secret     string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func New(secret string, accessTTL, refreshTTL time.Duration) (JWTToken, error) {
	if secret == "" {
		return nil, ErrMissingSecret
	}
	return &jwtToken{
		secret:     secret,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}, nil
}

func (s *jwtToken) GenerateAccessToken(user *domain.User) (string, error) {
	return s.generateToken(user.ID.String(), s.accessTTL)
}

func (s *jwtToken) GenerateRefreshToken(user *domain.User) (string, error) {
	return s.generateToken(user.ID.String(), s.refreshTTL)
}

func (s *jwtToken) ValidateToken(token string) (bool, error) {
	claims, err := s.parseToken(token)
	if err != nil {
		return false, err
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return false, ErrAccessTokenExpired
	}

	return true, nil
}

func (s *jwtToken) RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := s.parseToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to parse refresh token: %w", err)
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return "", ErrRefreshTokenExpired
	}

	return s.generateToken(claims.Subject, s.accessTTL)
}

func (s *jwtToken) generateToken(userID string, ttl time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    userID,
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *jwtToken) parseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	if claims.Subject == "" {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
