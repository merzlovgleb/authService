package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/itpark/market/auth/internal/domain"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func GenerateTokenPair(user domain.User, appID uuid.UUID, secret string) (*TokenPair, error) {
	accessToken, err := generateToken(user, appID, secret, 15*time.Minute, "access")
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateToken(user, appID, secret, 7*24*time.Hour, "refresh")
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateToken(user domain.User, appID uuid.UUID, secret string, duration time.Duration, tokenType string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID.String()
	claims["email"] = user.Email
	claims["app_id"] = appID.String()
	claims["type"] = tokenType
	claims["exp"] = time.Now().Add(duration).Unix()

	return token.SignedString([]byte(secret))
}
